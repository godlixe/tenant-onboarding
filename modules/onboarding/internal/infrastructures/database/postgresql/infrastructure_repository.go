package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
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
) ([]entities.Infrastructure, error) {
	var infrastructures []entities.Infrastructure
	tx := q.db.Model(&entities.Infrastructure{}).
		Where("product_id = ?", productID.String()).
		Where("user_count < user_limit").
		Where("deployment_model = ?", "pool").
		Order("user_count ASC").
		Find(&infrastructures)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return infrastructures, nil
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
