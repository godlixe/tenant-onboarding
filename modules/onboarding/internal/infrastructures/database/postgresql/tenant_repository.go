package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(
	db *gorm.DB,
) *TenantRepository {
	return &TenantRepository{
		db: db,
	}
}

func (r *TenantRepository) Create(ctx context.Context, tenant *entities.Tenant) error {
	tx := r.db.Model(&entities.Tenant{}).Create(tenant)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *TenantRepository) GetByID(
	ctx context.Context,
	tenantID valueobjects.TenantID,
) (*entities.Tenant, error) {
	var tenant *entities.Tenant
	tx := r.db.Model(&entities.Tenant{}).Limit(1).Find(&tenant)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tenant, nil
}
