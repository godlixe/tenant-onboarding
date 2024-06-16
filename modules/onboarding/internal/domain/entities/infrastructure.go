package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
	"tenant-onboarding/pkg/deployer/types"
)

type Infrastructure struct {
	ID              vo.InfrastructureID
	ProductID       vo.ProductID
	DeploymentModel types.DeploymentModel
	UserLimit       int
	Metadata        string
	database.Timestamp
}

func NewInfrastructure(
	id vo.InfrastructureID,
	productID vo.ProductID,
	name string,
	DeploymentModel types.DeploymentModel,
	userLimit int,
	metadata string,
) *Infrastructure {
	return &Infrastructure{
		ID:              id,
		ProductID:       productID,
		DeploymentModel: DeploymentModel,
		UserLimit:       userLimit,
		Metadata:        metadata,
	}
}
