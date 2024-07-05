package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type InfrastructureRepository struct {
	db *gorm.DB
}

func NewInfrastructureRepository(
	db *gorm.DB,
) *InfrastructureRepository {
	return &InfrastructureRepository{
		db: db,
	}
}

func (q *InfrastructureRepository) GetByProductIDInfraTypeOrdered(
	ctx context.Context,
	productID vo.ProductID,
) (*entities.Infrastructure, error) {
	var infrastructure entities.Infrastructure

	tx := q.db.Debug().Model(&entities.Infrastructure{}).
		Where("product_id = ?", productID.String()).
		Where("deployment_model = ?", "pool").
		Where(`(
		SELECT COUNT(tenants.id) 
		FROM tenants 
		WHERE tenants.infrastructure_id = infrastructures.id
	) < user_limit`).
		Order(`(
		SELECT COUNT(tenants.id) 
		FROM tenants 
		WHERE tenants.infrastructure_id = infrastructures.id
	) ASC`).
		Limit(1).
		Scan(&infrastructure)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &infrastructure, nil
}

func (q *InfrastructureRepository) Create(
	ctx context.Context,
	infrastructure *entities.Infrastructure,
) error {
	tx := q.db.Table("infrastructures").Create(infrastructure)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (q *InfrastructureRepository) Update(
	ctx context.Context,
	infrastructure *entities.Infrastructure,
) error {
	tx := q.db.Table("infrastructures").Updates(infrastructure)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (q *InfrastructureRepository) IncrementUser(
	ctx context.Context,
	id valueobjects.InfrastructureID,
) error {
	tx := q.db.Table("infrastructures").
		Where("id = ?", id.ID).
		Update("user_count", gorm.Expr("user_count + ?", 1))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
