package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type InfrastructureQuery struct {
	db *gorm.DB
}

func NewInfrastructureQuery(
	db *gorm.DB,
) *InfrastructureQuery {
	return &InfrastructureQuery{
		db: db,
	}
}

func (q *InfrastructureQuery) GetByProductIDInfraTypeOrdered(
	ctx context.Context,
	productID vo.ProductID,
) ([]entities.Infrastructure, error) {
	var (
		infrastructures []entities.Infrastructure
	)

	tx := q.db.Table("infrastructures").
		Where("product_id = ?", productID.String()).
		Where("user_count < user_limit").
		Where("deployment_model = pool").
		Order("user_count ASC").
		Distinct("name").
		Find(&infrastructures)
	if tx.Error != nil {
		return infrastructures, tx.Error
	}

	return infrastructures, nil
}
