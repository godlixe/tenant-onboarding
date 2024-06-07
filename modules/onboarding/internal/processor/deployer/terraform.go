package deployer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"time"

	"github.com/hashicorp/terraform-exec/tfexec"
)

type infraMetadata struct {
	Variables []map[string]string `json:"variables"`
}

func parseInfrastructures(
	infrastructureDefinition []InfrastructureDefinition,
	availableInfrastructures []entities.Infrastructure,
) (map[string]InfrastructureBlueprint, error) {
	infrastructureBlueprintMap := make(map[string]InfrastructureBlueprint)
	for _, data := range infrastructureDefinition {
		for key, val := range data {

			infrastructureBlueprintMap[key] = InfrastructureBlueprint{
				InfraStructureType: key,

				// defaults to create new, will be set
				// accordingly below.
				IsCreateNew:     true,
				DeploymentModel: val,
			}
		}
	}

	for availInfraKey, availInfra := range availableInfrastructures {
		val, ok := infrastructureBlueprintMap[availInfra.Name]
		if ok {
			var inframetadataJSON infraMetadata
			err := json.Unmarshal([]byte(availInfra.Metadata), &inframetadataJSON)
			if err != nil {
				return nil, err
			}

			infraVariables := make(map[string]string)
			for _, infraVarsIter := range inframetadataJSON.Variables {
				for infraVarkey, infraVarData := range infraVarsIter {
					infraVariables[infraVarkey] = infraVarData
				}
			}

			val.SetIsCreateNew(false)
			val.Variables = infraVariables
			val.InfrastructureEntity = &availableInfrastructures[availInfraKey]
			infrastructureBlueprintMap[availInfra.Name] = val
		}
	}

	return infrastructureBlueprintMap, nil
}

func createTenantExecutionDir(tenantID string) (string, error) {
	// executionPath is the path used by the pipeline to deploy
	var executionPath string = "/tmp"

	tenantDirPath := filepath.Join(
		executionPath,
		tenantID,
	)

	// create dir for new tenant
	if _, err := os.Stat(tenantID); os.IsNotExist(err) {
		err := os.Mkdir(tenantDirPath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return tenantDirPath, nil
}

func copyDeploymentRepository(schema DeploymentSchema, tenantDirPath string) error {
	gitCmd := exec.Command("/usr/bin/git", "clone", schema.DeploymentRepositoryURL)
	gitCmd.Dir = tenantDirPath
	gitCmd.Stderr = os.Stdout
	gitCmd.Stdout = os.Stdout

	err := gitCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func runTerraform(
	ctx context.Context,
	cfg Config,
	tenantDirPath string,
	tenantID string,
	deploymentSchema DeploymentSchema,
	infrastructureBlueprintMap map[string]InfrastructureBlueprint,
) error {
	tfEntryPoint := path.Join(
		tenantDirPath,
		deploymentSchema.TerraformExecutionPath,
	)

	tf, err := tfexec.NewTerraform(tfEntryPoint, cfg.TerraformExecPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tf.Init(ctx,
		tfexec.BackendConfig(fmt.Sprintf("prefix=%v", tenantID)),
		tfexec.BackendConfig(fmt.Sprintf("bucket=%v", cfg.TerraformBackendBucket)),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tf.SetStdout(os.Stdout)
	tfVariables := []tfexec.ApplyOption{
		tfexec.Var(fmt.Sprintf("project_id=%v", cfg.GoogleProjectID)),
		tfexec.Var(fmt.Sprintf("region=%v", cfg.GoogleDeploymentRegion)),
		tfexec.Var(fmt.Sprintf("tenant_name=%v", tenantID)),
		tfexec.Var(fmt.Sprintf("tenant_subdomain=%v", tenantID)),
		tfexec.Var(fmt.Sprintf("tenant_password=%v", tenantID)),
	}

	for key, val := range infrastructureBlueprintMap {
		// skip if not creating a new resource
		if val.IsCreateNew {
			tfVariables = append(tfVariables, tfexec.Var(fmt.Sprintf("is_create_%v=true", key)))
			continue
		}

		tfVariables = append(tfVariables, tfexec.Var(fmt.Sprintf("is_create_%v=false", key)))

		for infraBpKey, infraBpVal := range val.Variables {
			tfVariables = append(tfVariables, tfexec.Var(fmt.Sprintf("%v=%v", infraBpKey, infraBpVal)))
		}
	}

	err = tf.Apply(ctx, tfVariables...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tf.StatePull(ctx)
	if err != nil {
		return err
	}

	var out map[string]tfexec.OutputMeta
	out, err = tf.Output(ctx)
	if err != nil {
		return err
	}

	// poll until output is available
	for len(out) == 0 {
		time.Sleep(1000)
		out, err = tf.Output(ctx)
		if err != nil {
			return err
		}
		fmt.Println("polling")
	}

	for key, data := range out {
		keysSubstr := strings.Split(key, "_")
		val, ok := infrastructureBlueprintMap[keysSubstr[0]]
		if ok {
			var metadataMap map[string]string = make(map[string]string)
			metadataMap[key] = strings.ReplaceAll(string(data.Value), "\"", "")
			val.Metadata = append(val.Metadata, metadataMap)
			infrastructureBlueprintMap[keysSubstr[0]] = val
		}
	}

	return nil
}

func Deploy(
	ctx context.Context,
	cfg Config,
	tenantJob *entities.TenantDeploymentJob,
	infrastructureRepository repositories.InfrastructureRepository,
	tenantsInfrastructureRepository repositories.TenantsInfrastructuresRepository,
) error {
	var err error

	tenantDirPath, err := createTenantExecutionDir(tenantJob.TenantData.ID)
	if err != nil {
		return err
	}

	var deploymentSchema DeploymentSchema
	err = json.Unmarshal([]byte(tenantJob.ProductData.DeploymentSchema), &deploymentSchema)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = copyDeploymentRepository(deploymentSchema, tenantDirPath)
	if err != nil {
		return err
	}

	productIDValueObj, err := valueobjects.NewProductID(tenantJob.ProductData.ID)
	if err != nil {
		return err
	}
	availableInfrastructures, err := infrastructureRepository.GetByProductIDInfraTypeOrdered(ctx, productIDValueObj)
	if err != nil {
		return err
	}

	infrastructureBlueprintMap, err := parseInfrastructures(
		deploymentSchema.InfrastructureDefinition,
		availableInfrastructures,
	)
	if err != nil {
		return err
	}

	err = runTerraform(
		ctx,
		cfg,
		tenantDirPath,
		tenantJob.TenantData.ID,
		deploymentSchema,
		infrastructureBlueprintMap,
	)
	if err != nil {
		return err
	}

	insertInfrastructures(
		ctx,
		tenantJob.ProductData.ID,
		tenantJob.TenantData.ID,
		infrastructureBlueprintMap,
		infrastructureRepository,
		tenantsInfrastructureRepository,
	)

	return nil
}

func insertInfrastructures(
	ctx context.Context,
	productID string,
	tenantID string,
	infrastructureBlueprintMap map[string]InfrastructureBlueprint,
	infrastructureRepository repositories.InfrastructureRepository,
	tenantsInfrastructureRepository repositories.TenantsInfrastructuresRepository,
) error {
	productIDValueObj, err := valueobjects.NewProductID(productID)
	if err != nil {
		return err
	}

	tenantIDValueObj, err := valueobjects.NewTenantID(tenantID)
	if err != nil {
		return err
	}

	for _, data := range infrastructureBlueprintMap {
		if !data.IsCreateNew {
			err = infrastructureRepository.IncrementUser(ctx, data.InfrastructureEntity.ID)
			if err != nil {
				return err
			}

			err = tenantsInfrastructureRepository.Create(
				ctx,
				&entities.TenantsInfrastructures{
					TenantID:         tenantIDValueObj,
					InfrastructureID: data.InfrastructureEntity.ID,
				},
			)
			if err != nil {
				return err
			}

			continue
		}

		var metadata infraMetadata

		metadata.Variables = data.Metadata
		metadataJSONB, err := json.Marshal(metadata)
		if err != nil {
			return err
		}

		infrastructure := entities.Infrastructure{
			ID:              valueobjects.GenerateInfrastructureID(),
			ProductID:       productIDValueObj,
			Name:            data.InfraStructureType,
			DeploymentModel: data.DeploymentModel,
			UserCount:       1,
			UserLimit:       3,
			Metadata:        string(metadataJSONB),
		}

		err = infrastructureRepository.Create(ctx, &infrastructure)
		if err != nil {
			return err
		}

		err = tenantsInfrastructureRepository.Create(
			ctx,
			&entities.TenantsInfrastructures{
				TenantID:         tenantIDValueObj,
				InfrastructureID: infrastructure.ID,
			},
		)
		if err != nil {
			return err
		}

	}

	return nil
}
