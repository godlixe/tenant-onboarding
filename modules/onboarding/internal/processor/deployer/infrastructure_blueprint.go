package deployer

import (
	"errors"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

type InfrastructureBluePrint struct {
	InfraStructureType string
	DeploymentModel    valueobjects.DeploymentModel
	Metadata           []map[string]string
	IsCreateNew        bool
	Variables          map[string]string
}

func (i *InfrastructureBluePrint) SetIsCreateNew(isCreateNew bool) error {
	if i.DeploymentModel == valueobjects.Silo && !isCreateNew {
		return errors.New("invalid parameter for silo objects")
	}

	i.IsCreateNew = isCreateNew
	return nil
}
