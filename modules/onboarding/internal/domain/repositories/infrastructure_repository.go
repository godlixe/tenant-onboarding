package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

type InfrastructureRepository interface {
	GetByProductIDInfraTypeOrdered(
		ctx context.Context,
		productID vo.ProductID,
	) ([]entities.Infrastructure, error)

	Create(
		ctx context.Context,
		infrastructure *entities.Infrastructure,
	) error

	Update(
		ctx context.Context,
		infrastructure *entities.Infrastructure,
	) error
}
