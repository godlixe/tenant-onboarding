package pipeline

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"tenant-onboarding/modules/onboarding/internal/domain/entities"
// 	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
// 	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

// 	"github.com/google/uuid"
// )

// // func createInfrastructureDir(tenantID string) (string, error) {
// // 	// executionPath is the path used by the pipeline to deploy
// // 	var executionPath string = "/tmp"

// // 	tenantDirPath := filepath.Join(
// // 		executionPath,
// // 		tenantID,
// // 	)

// // 	// create dir for new tenant
// // 	if _, err := os.Stat(tenantDirPath); os.IsNotExist(err) {
// // 		err := os.Mkdir(tenantDirPath, fs.ModePerm)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 			return "", err
// // 		}
// // 	}

// // 	return tenantDirPath, nil
// // }

// // func isValidURL(repositoryURL string) (bool, error) {
// // 	_, err := url.Parse(repositoryURL)
// // 	if err != nil {
// // 		return false, errorx.ErrInvalidRepositoryURL
// // 	}

// // 	return true, nil
// // }

// // func copyDeploymentRepository(repositoryURL, tenantDirPath string) error {
// // 	if validURL, err := isValidURL(repositoryURL); !validURL {
// // 		return err
// // 	}
// // 	cmd := exec.Command("/usr/bin/git", "clone", repositoryURL)
// // 	cmd.Dir = tenantDirPath
// // 	cmd.Stderr = os.Stdout
// // 	cmd.Stdout = os.Stdout

// // 	err := cmd.Run()
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // func runTerraform(
// // 	ctx context.Context,
// // 	cfg Config,
// // 	terraformDirPath string,
// // 	infrastructureID string,
// // 	deploymentSchema DeploymentSchema,
// // ) (string, error) {
// // 	tfEntryPoint := path.Join(
// // 		terraformDirPath,
// // 		deploymentSchema.TerraformExecutionPath,
// // 	)

// // 	tf, err := tfexec.NewTerraform(tfEntryPoint, cfg.TerraformExecPath)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return "", err
// // 	}

// // 	err = tf.Init(ctx,
// // 		tfexec.BackendConfig(fmt.Sprintf("prefix=%v", infrastructureID)),
// // 		tfexec.BackendConfig(fmt.Sprintf("bucket=%v", cfg.TerraformBackendBucket)),
// // 	)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return "", err
// // 	}

// // 	tf.SetStdout(os.Stdout)
// // 	tfVariables := []tfexec.ApplyOption{
// // 		tfexec.Var(fmt.Sprintf("project_id=%v", cfg.GoogleProjectID)),
// // 		tfexec.Var(fmt.Sprintf("region=%v", cfg.GoogleDeploymentRegion)),
// // 		tfexec.Var(fmt.Sprintf("infrastructure_id=%v", infrastructureID)),
// // 	}

// // 	err = tf.Apply(ctx, tfVariables...)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return "", err
// // 	}

// // 	_, err = tf.StatePull(ctx)
// // 	if err != nil {
// // 		return "", err
// // 	}

// // 	var out map[string]tfexec.OutputMeta
// // 	out, err = tf.Output(ctx)
// // 	if err != nil {
// // 		return "", err
// // 	}

// // 	// poll until output is available
// // 	for len(out) == 0 {
// // 		time.Sleep(1000)
// // 		out, err = tf.Output(ctx)
// // 		if err != nil {
// // 			return "", err
// // 		}
// // 		fmt.Println("polling")
// // 	}

// // 	var metadata string
// // 	for key, data := range out {
// // 		if key == "metadata" {
// // 			metadata = string(data.Value)
// // 		}
// // 	}

// // 	return metadata, nil
// // }

// func initiateInfrastructure(
// 	ctx context.Context,
// 	cfg Config,
// 	deploymentSchema DeploymentSchema,
// ) (*RawInfrastructure, error) {
// 	var err error
// 	infrastructureDirID := uuid.NewString()
// 	infrastructureDirPath, err := createInfrastructureDir(infrastructureDirID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = copyDeploymentRepository(
// 		deploymentSchema.DeploymentRepositoryURL,
// 		infrastructureDirPath,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	infrastructureMetadata, err := runTerraform(
// 		ctx,
// 		cfg,
// 		infrastructureDirPath,
// 		infrastructureDirID,
// 		deploymentSchema,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &RawInfrastructure{
// 		ID:          infrastructureDirID,
// 		Metadata:    infrastructureMetadata,
// 		IsCreateNew: true,
// 	}, nil
// }

// func Deploy(
// 	ctx context.Context,
// 	cfg Config,
// 	tenantJob *entities.TenantDeploymentJob,
// 	infrastructureRepository repositories.InfrastructureRepository,
// 	tenantsInfrastructuresRepository repositories.TenantsInfrastructuresRepository,
// 	tenantRepository repositories.TenantRepository,
// 	productRepository repositories.ProductRepository,
// ) error {
// 	var err error

// 	productIDValueObj, err := valueobjects.NewProductID(tenantJob.ProductID)
// 	if err != nil {
// 		return err
// 	}

// 	product, err := productRepository.GetByID(ctx, productIDValueObj)
// 	if err != nil {
// 		return err
// 	}

// 	availableInfrastructure, err := infrastructureRepository.GetByProductIDInfraTypeOrdered(ctx, productIDValueObj)
// 	if err != nil {
// 		return err
// 	}

// 	var deploymentSchema DeploymentSchema
// 	err = json.Unmarshal([]byte(product.DeploymentSchema), &deploymentSchema)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	var infrastructureMetadata *RawInfrastructure = &RawInfrastructure{}
// 	fmt.Println(availableInfrastructure)
// 	if (*availableInfrastructure == entities.Infrastructure{}) ||
// 		(deploymentSchema.DeploymentModel == valueobjects.Silo) {
// 		infrastructureMetadata, err = initiateInfrastructure(
// 			ctx,
// 			cfg,
// 			deploymentSchema,
// 		)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		infrastructureMetadata.ID = availableInfrastructure.ID.String()
// 		infrastructureMetadata.IsCreateNew = false
// 	}

// 	fmt.Println(infrastructureMetadata)

// 	err = processInfrastructures(
// 		ctx,
// 		tenantJob.ProductID,
// 		tenantJob.TenantID,
// 		infrastructureMetadata,
// 		deploymentSchema,
// 		infrastructureRepository,
// 		tenantRepository,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func processInfrastructures(
// 	ctx context.Context,
// 	productID string,
// 	tenantID string,
// 	rawInfrastructure *RawInfrastructure,
// 	deploymentSchema DeploymentSchema,
// 	infrastructureRepository repositories.InfrastructureRepository,
// 	tenantRepository repositories.TenantRepository,
// ) error {
// 	productIDValueObj, err := valueobjects.NewProductID(productID)
// 	if err != nil {
// 		return err
// 	}

// 	tenantIDValueObj, err := valueobjects.NewTenantID(tenantID)
// 	if err != nil {
// 		return err
// 	}

// 	infrastructureID, err := valueobjects.NewInfrastructureID(rawInfrastructure.ID)
// 	if err != nil {
// 		return err
// 	}

// 	if rawInfrastructure.IsCreateNew {
// 		var limit int = 1

// 		if deploymentSchema.DeploymentModel == valueobjects.Pool {
// 			limit = 3
// 		}

// 		err = infrastructureRepository.Create(ctx, &entities.Infrastructure{
// 			ID:              infrastructureID,
// 			ProductID:       productIDValueObj,
// 			DeploymentModel: deploymentSchema.DeploymentModel,
// 			UserCount:       1,
// 			UserLimit:       limit,
// 			Metadata:        rawInfrastructure.Metadata,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		err = infrastructureRepository.IncrementUser(ctx, infrastructureID)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err = tenantRepository.Update(
// 		ctx,
// 		&entities.Tenant{
// 			ID:               tenantIDValueObj,
// 			InfrastructureID: infrastructureID,
// 		},
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
