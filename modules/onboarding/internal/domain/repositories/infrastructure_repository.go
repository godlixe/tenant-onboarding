package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
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

	IncrementUser(
		ctx context.Context,
		id valueobjects.InfrastructureID,
	) error
}
