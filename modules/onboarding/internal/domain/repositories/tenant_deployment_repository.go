package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
)

type TenantDeploymentRepository interface {
	PublishTenantDeploymentJob(context.Context, *entities.TenantDeploymentJob) error
}
