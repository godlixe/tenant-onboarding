package services

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"
)

type TenantService interface {
	CreateTenant(
		ctx context.Context,
		tenant *entity.Tenant,
	) (*entity.Tenant, error)
}
