package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/app/queries"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"

	"gorm.io/gorm"
)

type TenantQuery struct {
	db *gorm.DB
}

func NewTenantQuery(
	db *gorm.DB,
) *TenantQuery {
	return &TenantQuery{
		db: db,
	}
}

func (q *TenantQuery) GetByID(ctx context.Context, id string) (queries.Tenant, error) {
	var (
		tenant queries.Tenant
	)

	tx := q.db.Table("tenants").Where("id = ?", id).Find(&tenant)
	if tx.Error != nil {
		return tenant, tx.Error
	}

	return tenant, nil
}

func (q *TenantQuery) GetAllByOrganizationID(ctx context.Context, organizationID string) ([]queries.Tenant, error) {
	var (
		tenants []queries.Tenant
	)

	tx := q.db.Model(&entities.Tenant{}).
		Where("organization_id = ?", organizationID).
		Preload("Product.App").
		Find(&tenants)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tenants, nil
}
