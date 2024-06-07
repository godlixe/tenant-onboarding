package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"

	"gorm.io/gorm"
)

type TenantsInfrastructuresRepository struct {
	db *gorm.DB
}

func NewTenantsInfrastructuresRepository(
	db *gorm.DB,
) *TenantsInfrastructuresRepository {
	return &TenantsInfrastructuresRepository{
		db: db,
	}
}

func (t *TenantsInfrastructuresRepository) Create(
	ctx context.Context,
	tenantsInfrastructures *entities.TenantsInfrastructures,
) error {
	tx := t.db.Model(&entities.TenantsInfrastructures{}).Create(tenantsInfrastructures)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
