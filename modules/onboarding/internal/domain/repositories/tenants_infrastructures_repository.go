package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
)

type TenantsInfrastructuresRepository interface {
	Create(
		ctx context.Context,
		tenantsInfrastructures *entities.TenantsInfrastructures,
	) error
}
