package postgresql

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{
		db: db,
	}
}

func (r *TenantRepository) FindById(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Tenant, error) {
	var tenant entity.Tenant

	tx := r.db.Model(entity.Tenant{}).Where("id = ?", id).First(&tenant)

	return &tenant, tx.Error
}

func (r *TenantRepository) CreateTenant(
	ctx context.Context,
	tenant *entity.Tenant,
) error {
	tx := r.db.Model(entity.Tenant{}).Create(&tenant)

	return tx.Error
}

func (r *TenantRepository) CreateUserTenant(
	ctx context.Context,
	userTenant *entity.UserTenant,
) error {
	tx := r.db.Model(entity.UserTenant{}).Create(&userTenant)

	return tx.Error
}

func (r *TenantRepository) UpdateTenant(
	ctx context.Context,
	tenant *entity.Tenant,
) error {
	tx := r.db.Model(entity.Tenant{}).Updates(&tenant)

	return tx.Error
}
