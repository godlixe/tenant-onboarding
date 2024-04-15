package repository

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"

	"github.com/google/uuid"
)

type TenantRepository interface {
	FindById(
		ctx context.Context,
		id uuid.UUID,
	) (*entity.Tenant, error)

	CreateTenant(
		ctx context.Context,
		tenant *entity.Tenant,
	) error

	UpdateTenant(
		ctx context.Context,
		tenant *entity.Tenant,
	) error
}
