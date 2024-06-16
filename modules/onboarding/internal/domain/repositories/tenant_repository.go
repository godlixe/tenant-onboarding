package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entities.Tenant) error
	Update(ctx context.Context, tenant *entities.Tenant) error
	GetByID(ctx context.Context, id valueobjects.TenantID) (*entities.Tenant, error)
}
