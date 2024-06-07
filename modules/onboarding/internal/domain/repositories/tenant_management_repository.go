package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
)

type TenantManagementRepository interface {
	CreateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error)
}
