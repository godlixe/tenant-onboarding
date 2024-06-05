package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/app/queries"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

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

func (q *TenantQuery) GetByID(ctx context.Context, id vo.ProductID) (queries.Tenant, error) {
	var (
		tenant queries.Tenant
	)

	tx := q.db.Table("tenants").Where("id = ?", id.String()).Find(&tenant)
	if tx.Error != nil {
		return tenant, tx.Error
	}

	return tenant, nil
}
