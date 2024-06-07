package deployer

import (
	"errors"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

// DeploymentSchema defines the deployment schema of
// a product. A deployment schema of a product is predefined.
type DeploymentSchema struct {
	// Github repository of the directory for deployment.
	DeploymentRepositoryURL string `json:"deployment_repository_url"`
	// Terraform execution path.
	TerraformExecutionPath string `json:"terraform_execution_path"`
	// Custom script execution path.
	ScriptExecutionPath string `json:"script_execution_path"`

	InfrastructureDefinition []InfrastructureDefinition `json:"infrastructure_blueprint"`
}

type InfrastructureDefinition map[string]valueobjects.DeploymentModel

type InfrastructureBlueprint struct {
	InfraStructureType   string
	DeploymentModel      valueobjects.DeploymentModel
	Metadata             []map[string]string
	IsCreateNew          bool
	Variables            map[string]string
	InfrastructureEntity *entities.Infrastructure
}

func (i *InfrastructureBlueprint) SetIsCreateNew(isCreateNew bool) error {
	if i.DeploymentModel == valueobjects.Silo && !isCreateNew {
		return errors.New("invalid parameter for silo objects")
	}

	i.IsCreateNew = isCreateNew
	return nil
}
