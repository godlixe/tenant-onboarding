package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

type ProductRepository interface {
	GetByID(ctx context.Context, id vo.ProductID) (*entities.Product, error)
}
