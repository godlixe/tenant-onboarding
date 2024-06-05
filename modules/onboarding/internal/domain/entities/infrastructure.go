package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
)

type Infrastructure struct {
	ID              vo.InfrastructureID
	ProductID       vo.ProductID
	Name            string
	DeploymentModel vo.DeploymentModel
	UserCount       int
	UserLimit       int
	Metadata        string
	database.Timestamp
}

func NewInfrastructure(
	id vo.InfrastructureID,
	productID vo.ProductID,
	name string,
	DeploymentModel vo.DeploymentModel,
	userCount int,
	userLimit int,
	metadata string,
) *Infrastructure {
	return &Infrastructure{
		ID:              id,
		ProductID:       productID,
		Name:            name,
		DeploymentModel: DeploymentModel,
		UserCount:       userCount,
		UserLimit:       userLimit,
		Metadata:        metadata,
	}
}
