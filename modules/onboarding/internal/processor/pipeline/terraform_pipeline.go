package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/modules/onboarding/internal/errorx"
	"tenant-onboarding/pkg/deployer"
	"tenant-onboarding/pkg/deployer/types"
	"tenant-onboarding/pkg/framework"
	"tenant-onboarding/providers"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/samber/do"
)

type TerraformDeployer struct {
	app       *providers.App
	cfg       deployer.Config
	dataStore map[string]any
}

func NewTerraformDeployer(
	app *providers.App,
	cfg deployer.Config,
) *TerraformDeployer {
	return &TerraformDeployer{
		app:       app,
		cfg:       cfg,
		dataStore: make(map[string]any),
	}
}

func (t *TerraformDeployer) GetData(
	ctx context.Context,
	tenantDeploymentJob types.TenantDeploymentJob,
) (*types.DeploymentSchema, error) {
	var err error

	tenantIDValueObj, err := valueobjects.NewTenantID(tenantDeploymentJob.TenantID)
	if err != nil {
		return nil, err
	}

	tenantRepository, err := do.Invoke[repositories.TenantRepository](t.app.Injector)
	if err != nil {
		return nil, err
	}

	tenant, err := tenantRepository.GetByID(ctx, tenantIDValueObj)
	if err != nil {
		return nil, err
	}

	productIDValueObj, err := valueobjects.NewProductID(tenant.ProductID.String())
	if err != nil {
		return nil, err
	}

	productRepository, err := do.Invoke[repositories.ProductRepository](t.app.Injector)
	if err != nil {
		return nil, err
	}

	product, err := productRepository.GetByID(ctx, productIDValueObj)
	if err != nil {
		return nil, err
	}
	if tenant.Status != valueobjects.TenantCreated {
		fmt.Println("tenant is onboarded already", tenant.Status)
		return nil, errors.New("tenant is already onboarded")
	}

	// save data to data store
	t.dataStore["organization_id"] = tenant.OrganizationID.String()
	t.dataStore["tenant_id"] = tenant.ID.String()
	t.dataStore["product_id"] = productIDValueObj.String()
	t.dataStore["app_id"] = product.AppID

	var deploymentSchema types.DeploymentSchema
	err = json.Unmarshal([]byte(product.DeploymentSchema), &deploymentSchema)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	deploymentSchema.ProductID = productIDValueObj.String()

	return &deploymentSchema, nil
}

func (t *TerraformDeployer) Initiate(
	ctx context.Context,
	deploymentSchema *types.DeploymentSchema,
) (*types.RawInfrastructure, error) {
	infrastructureRepository, err := do.Invoke[repositories.InfrastructureRepository](t.app.Injector)
	if err != nil {
		return nil, err
	}
	productIDValueObj, err := valueobjects.NewProductID(deploymentSchema.ProductID)
	if err != nil {
		return nil, err
	}

	if deploymentSchema.DeploymentModel == types.Pool {
		// get available infrastructures
		availableInfrastructure, err := infrastructureRepository.GetByProductIDInfraTypeOrdered(ctx, productIDValueObj)
		if err != nil {
			return nil, err
		}

		if (*availableInfrastructure != entities.Infrastructure{}) {
			type MetadataHolder struct {
				Metadata            json.RawMessage `json:"metadata"`
				ResourceInformation json.RawMessage `json:"resource_information"`
			}

			var availInfraMetadata MetadataHolder

			err = json.Unmarshal(
				[]byte(availableInfrastructure.Metadata),
				&availInfraMetadata,
			)
			if err != nil {
				return nil, err
			}

			return &types.RawInfrastructure{
				ID: availableInfrastructure.ID.String(),
				Metadata: &types.InfraOutput{
					Metadata:             json.RawMessage(availableInfrastructure.Metadata),
					ResourceInformations: availInfraMetadata.ResourceInformation,
				},
				IsCreateNew: false,
			}, nil
		}

	}
	// if infrastructure is available
	infrastructureID := uuid.NewString()

	return &types.RawInfrastructure{
		ID:          infrastructureID,
		IsCreateNew: true,
	}, nil

}

func createInfrastructureDir(infrastructureID string) (string, error) {
	// executionPath is the path used by the pipeline to deploy
	var executionPath string = "/tmp"

	deploymentDirPath := filepath.Join(
		executionPath,
		infrastructureID,
	)

	// create dir for new tenant
	if _, err := os.Stat(deploymentDirPath); os.IsNotExist(err) {
		err := os.Mkdir(deploymentDirPath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return deploymentDirPath, nil
}

func isValidURL(repositoryURL string) (bool, error) {
	_, err := url.Parse(repositoryURL)
	if err != nil {
		return false, errorx.ErrInvalidRepositoryURL
	}

	return true, nil
}

func copyDeploymentRepository(repositoryURL, deploymentDirPath string) error {
	if validURL, err := isValidURL(repositoryURL); !validURL {
		return err
	}

	cmd := exec.Command("/usr/bin/git", "clone", repositoryURL)

	cmd.Dir = deploymentDirPath

	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func runTerraform(
	ctx context.Context,
	cfg deployer.Config,
	deploymentSchema *types.DeploymentSchema,
	rawInfrastructure *types.RawInfrastructure,
) (*types.InfraOutput, error) {
	deploymentDirPath, err := createInfrastructureDir(rawInfrastructure.ID)
	if err != nil {
		return nil, err
	}

	err = copyDeploymentRepository(
		deploymentSchema.DeploymentRepositoryURL,
		deploymentDirPath,
	)
	if err != nil {
		return nil, err
	}

	tfEntryPoint := path.Join(
		deploymentDirPath,
		deploymentSchema.TerraformExecutionPath,
	)

	tf, err := tfexec.NewTerraform(tfEntryPoint, cfg.TerraformExecPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = tf.Init(ctx,
		tfexec.BackendConfig(fmt.Sprintf("prefix=%v", rawInfrastructure.ID)),
		tfexec.BackendConfig(fmt.Sprintf("bucket=%v", cfg.TerraformBackendBucket)),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tf.SetStdout(os.Stdout)
	tfVariables := []tfexec.ApplyOption{
		tfexec.Var(fmt.Sprintf("provider_id=%v", cfg.GoogleProjectID)),
		tfexec.Var(fmt.Sprintf("infrastructure_id=%v", rawInfrastructure.ID)),
	}

	err = tf.Apply(ctx, tfVariables...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = tf.StatePull(ctx)
	if err != nil {
		return nil, err
	}

	var out map[string]tfexec.OutputMeta
	out, err = tf.Output(ctx)
	if err != nil {
		return nil, err
	}

	// poll until output is available
	for len(out) == 0 {
		time.Sleep(1000)
		out, err = tf.Output(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Println("polling")
	}

	var metadata json.RawMessage
	var resourceInformation json.RawMessage
	for key, data := range out {
		if key == "metadata" {
			// metadataBuffer := new(bytes.Buffer)

			// err = json.Compact(metadataBuffer, data.Value)
			// if err != nil {
			// 	return nil, err
			// }

			metadata = data.Value
		}

		if key == "resource_information" {
			// resourceInformationBuffer := new(bytes.Buffer)

			// err = json.Compact(resourceInformationBuffer, data.Value)
			// if err != nil {
			// 	return nil, err
			// }

			resourceInformation = data.Value
		}
	}

	return &types.InfraOutput{
		Metadata:             metadata,
		ResourceInformations: resourceInformation,
	}, nil
}

func (t *TerraformDeployer) Deploy(
	ctx context.Context,
	deploymentSchema *types.DeploymentSchema,
	rawInfrastructure *types.RawInfrastructure,
) (*types.RawInfrastructure, error) {
	if !rawInfrastructure.IsCreateNew {
		return rawInfrastructure, nil
	}

	metadata, err := runTerraform(
		ctx,
		t.cfg,
		deploymentSchema,
		rawInfrastructure,
	)
	if err != nil {
		return nil, err
	}

	rawInfrastructure.Metadata = metadata

	return rawInfrastructure, nil
}

func persistInfrastructure(
	ctx context.Context,
	app *providers.App,
	tenantDeploymentJob types.TenantDeploymentJob,
	deploymentSchema *types.DeploymentSchema,
	rawInfrastructure *types.RawInfrastructure,
) error {
	productIDValueObj, err := valueobjects.NewProductID(deploymentSchema.ProductID)
	if err != nil {
		return err
	}

	tenantIDValueObj, err := valueobjects.NewTenantID(tenantDeploymentJob.TenantID)
	if err != nil {
		return err
	}

	infrastructureID, err := valueobjects.NewInfrastructureID(rawInfrastructure.ID)
	if err != nil {
		return err
	}

	infrastructureRepository, err := do.Invoke[repositories.InfrastructureRepository](app.Injector)
	if err != nil {
		return err
	}

	tenantRepository, err := do.Invoke[repositories.TenantRepository](app.Injector)
	if err != nil {
		return err
	}

	if rawInfrastructure.IsCreateNew {
		var limit int = 1

		if deploymentSchema.DeploymentModel == types.Pool {
			limit = 3
		}

		metadataBytes, err := json.Marshal(rawInfrastructure.Metadata)
		if err != nil {
			return err
		}

		err = infrastructureRepository.Create(ctx, &entities.Infrastructure{
			ID:              infrastructureID,
			ProductID:       productIDValueObj,
			DeploymentModel: deploymentSchema.DeploymentModel,
			UserLimit:       limit,
			Metadata:        string(metadataBytes),
		})
		if err != nil {
			return err
		}
	}
	tenantStatusStr := "activated"

	if framework.CheckIntegratedMode() {
		tenantStatusStr = "created"
	}

	tenantStatus, err := valueobjects.NewTenantStatus(tenantStatusStr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tenantRepository.Update(ctx, &entities.Tenant{
		ID:                  tenantIDValueObj,
		InfrastructureID:    &infrastructureID,
		Status:              tenantStatus,
		ResourceInformation: string(rawInfrastructure.Metadata.ResourceInformations),
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TerraformDeployer) PostDeployment(
	ctx context.Context,
	tenantDeploymentJob types.TenantDeploymentJob,
	rawInfrastructure *types.RawInfrastructure,
	deploymentSchema *types.DeploymentSchema,
) error {
	err := persistInfrastructure(
		ctx,
		t.app,
		tenantDeploymentJob,
		deploymentSchema,
		rawInfrastructure,
	)
	if err != nil {
		return err
	}

	metadataBytes, err := json.Marshal(rawInfrastructure.Metadata)
	if err != nil {
		return err
	}

	deploymentDir := fmt.Sprintf("/tmp/%v", rawInfrastructure.ID)

	if rawInfrastructure.IsCreateNew {
		err = runInitScript(path.Join(deploymentDir, deploymentSchema.InitScriptPath), string(metadataBytes))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// send tenant onboarded event
	tenantOnboardedRepository, err := do.Invoke[repositories.TenantOnboardedRepository](t.app.Injector)
	if err != nil {
		return err
	}

	tenantOnboardedEvent := types.TenantOnboardedEvent{
		TenantID:            t.dataStore["tenant_id"].(string),
		OrgID:               t.dataStore["organization_id"].(string),
		AppID:               t.dataStore["app_id"].(int),
		ProductID:           deploymentSchema.ProductID,
		Metadata:            rawInfrastructure.Metadata.Metadata,
		ResourceInformation: rawInfrastructure.Metadata.ResourceInformations,
		Timestamp:           time.Now(),
	}

	err = tenantOnboardedRepository.PublishTenantOnboarded(
		ctx,
		&tenantOnboardedEvent,
	)
	if err != nil {
		return err
	}

	err = cleanup(deploymentDir)
	if err != nil {
		return err
	}

	return nil
}

func cleanup(
	infrastructureDirPath string,
) error {
	err := os.RemoveAll(infrastructureDirPath)
	if err != nil {
		return err
	}

	return err
}

func runInitScript(initScriptPath string, metadataJSONStr string) error {
	cmd := exec.Command(initScriptPath, fmt.Sprintf(`--metadata=%v`, metadataJSONStr))

	cmd.Dir = path.Dir(initScriptPath)

	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
