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
	tx := r.db.Model(&entities.Tenant{}).
		Where("id = ?", tenantID.String()).
		Limit(1).
		Find(&tenant)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tenant, nil
}

func (r *TenantRepository) GetByAppIDOrgID(
	ctx context.Context,
	appID valueobjects.AppID,
	organizationID valueobjects.OrganizationID,
) (*entities.Tenant, error) {
	var tenant *entities.Tenant
	tx := r.db.Debug().Model(&entities.Tenant{}).
		Joins("JOIN products ON products.id = tenants.product_id").
		Where("organization_id = ?", organizationID).
		Where("products.app_id = ?", appID.Int()).
		Limit(1).
		Find(&tenant)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tenant, nil
}

func (r *TenantRepository) Update(ctx context.Context, tenant *entities.Tenant) error {
	tx := r.db.Debug().Model(&entities.Tenant{}).
		Where("id = ?", tenant.ID).
		Updates(tenant)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
