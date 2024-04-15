package repository

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"
)

type TenantPublisher interface {
	Publish(
		ctx context.Context,
		tenant *entity.Tenant,
	) error
}
