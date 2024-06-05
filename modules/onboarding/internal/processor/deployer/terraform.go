package deployer

import (
	"context"
	"encoding/json"
	"errors"
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

	"github.com/hashicorp/terraform-exec/tfexec"
)

func Deploy(
	ctx context.Context,
	cfg Config,
	tenantJob *entities.TenantDeploymentJob,
	infrastructureRepository repositories.InfrastructureRepository,
) error {
	var err error

	// executionPath is the path used by the pipeline to deploy
	var executionPath string = "/tmp"

	tenantId := tenantJob.TenantData.ID
	tenantDirPath := filepath.Join(
		executionPath,
		tenantId,
	)

	// create dir for new tenant
	if _, err := os.Stat(tenantId); os.IsNotExist(err) {
		err := os.Mkdir(tenantDirPath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Get deployment schema for product.
	var deploymentSchema map[string]interface{}
	err = json.Unmarshal([]byte(tenantJob.ProductData.DeploymentSchema), &deploymentSchema)
	if err != nil {
		return err
	}

	// infrastructureMetadata stores infrastructure related metadata according to
	// types defined in `infraStructureBlueprintJson`.
	var infrastructureBlueprintMap map[string]InfrastructureBluePrint = make(map[string]InfrastructureBluePrint)
	var infraStructureBlueprintJson []any = deploymentSchema["infrastructure_blueprint"].([]any)
	for _, data := range infraStructureBlueprintJson {
		infraData := data.(map[string]any)
		for key, val := range infraData {
			deploymentModelString := val.(string)
			deploymentModel, err := valueobjects.NewDeploymentModel(deploymentModelString)
			if err != nil {
				return err
			}

			infrastructureBlueprintMap[key] = InfrastructureBluePrint{
				InfraStructureType: key,

				// defaults to create new, will be set
				// accordingly below.
				IsCreateNew:     true,
				DeploymentModel: deploymentModel,
			}
		}
	}
	productIDValueObj, err := valueobjects.NewProductID(tenantJob.ProductData.ID)
	if err != nil {
		return err
	}
	availableInfrastructures, err := infrastructureRepository.GetByProductIDInfraTypeOrdered(context.TODO(), productIDValueObj)
	if err != nil {
		return err
	}

	for _, data := range availableInfrastructures {
		val, ok := infrastructureBlueprintMap[data.Name]
		if ok {
			var inframetadataJSON map[string]any

			err = json.Unmarshal([]byte(data.Metadata), &inframetadataJSON)
			if err != nil {
				return err
			}

			var infraVariables map[string]string = make(map[string]string)
			infraMetadataVariables := inframetadataJSON["variables"].([]any)
			for _, data := range infraMetadataVariables {
				mp := data.(map[string]any)
				for mpKey, mpData := range mp {
					if _, ok := mpData.(string); !ok {
						return errors.New("invalid string")
					}
					infraVariables[mpKey] = mpData.(string)
				}
			}
			val.SetIsCreateNew(false)
			val.Variables = infraVariables
			infrastructureBlueprintMap[data.Name] = val
		}
	}

	// Clone repository to execution path.
	terraformRepoURL := string(deploymentSchema["terraform_repository_url"].(string))
	gitCmd := exec.Command("/usr/bin/git", "clone", terraformRepoURL)
	gitCmd.Dir = tenantDirPath
	gitCmd.Stderr = os.Stdout
	gitCmd.Stdout = os.Stdout

	err = gitCmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	tfEntryPoint := path.Join(tenantDirPath, string(
		deploymentSchema["terraform_entrypoint_dir"].(string)),
	)

	tf, err := tfexec.NewTerraform(tfEntryPoint, cfg.TerraformExecPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tf.Init(ctx,
		tfexec.BackendConfig(fmt.Sprintf("prefix=%s", tenantJob.TenantData.ID)),
		tfexec.BackendConfig("bucket=terraform-dep"),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tf.SetStdout(os.Stdout)
	tfVariables := []tfexec.ApplyOption{
		tfexec.Var("project_id=static-booster-418207"),
		tfexec.Var("region=asia-southeast2"),
		tfexec.Var(fmt.Sprintf("tenant_name=%v", tenantJob.TenantData.ID)),
		tfexec.Var(fmt.Sprintf("tenant_subdomain=%v", tenantJob.TenantData.ID)),
		tfexec.Var(fmt.Sprintf("tenant_password=%v", tenantJob.TenantData.ID)),
	}
	for key, val := range infrastructureBlueprintMap {
		// skip if not creating a new resource
		if !val.IsCreateNew {
			continue
		}

		tfVariables = append(tfVariables, tfexec.Var(fmt.Sprintf("is_create_%v=false", key)))

		for infraBpKey, infraBpVal := range val.Variables {
			tfVariables = append(tfVariables, tfexec.Var(fmt.Sprintf("%v=%v", infraBpKey, infraBpVal)))
		}
	}
	// for _, data := range tfVariables {
	// 	tmpData := data.(*tfexec.VarOption)
	// 	fmt.Println(tmpData)
	// }

	err = tf.Apply(ctx, tfVariables...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var out map[string]tfexec.OutputMeta
	out, err = tf.Output(ctx)
	if err != nil {
		return err
	}

	_, err = tf.StatePull(ctx)
	if err != nil {
		return err
	}

	// looping until output is available
	for len(out) == 0 {
		out, err = tf.Output(ctx)
		if err != nil {
			return err
		}
		fmt.Println("polling")
	}

	for key, data := range out {
		fmt.Println(key, string(data.Value))
		keysSubstr := strings.Split(key, "_")
		fmt.Println(keysSubstr)
		val, ok := infrastructureBlueprintMap[keysSubstr[0]]
		if ok {
			var metadataMap map[string]string = make(map[string]string)
			metadataMap[key] = strings.ReplaceAll(string(data.Value), "\"", "")
			val.Metadata = append(val.Metadata, metadataMap)

			infrastructureBlueprintMap[keysSubstr[0]] = val
		}
	}

	insertInfrastructures(
		tenantJob.ProductData.ID,
		infrastructureBlueprintMap,
		infrastructureRepository,
	)

	return nil
}

func insertInfrastructures(
	productID string,
	infrastructureBlueprintMap map[string]InfrastructureBluePrint,
	infrastructureRepository repositories.InfrastructureRepository,
) error {
	productIDValueObj, err := valueobjects.NewProductID(productID)
	if err != nil {
		return err
	}
	for _, data := range infrastructureBlueprintMap {
		if !data.IsCreateNew {
			continue
		}

		variablesMap := make(map[string][]map[string]string)
		variablesMap["variables"] = data.Metadata
		metadataJSONB, err := json.Marshal(variablesMap)
		if err != nil {
			return err
		}

		infrastructure := entities.Infrastructure{
			ID:              valueobjects.GenerateInfrastructureID(),
			ProductID:       productIDValueObj,
			Name:            data.InfraStructureType,
			DeploymentModel: data.DeploymentModel,
			UserCount:       1,
			UserLimit:       100,
			Metadata:        string(metadataJSONB),
		}

		infrastructureRepository.Create(context.TODO(), &infrastructure)
	}

	return nil
}
