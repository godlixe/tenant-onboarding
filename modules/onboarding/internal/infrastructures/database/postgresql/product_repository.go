package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(
	db *gorm.DB,
) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (q *ProductRepository) GetByID(
	ctx context.Context,
	productID valueobjects.ProductID,
) (*entities.Product, error) {
	var product entities.Product

	tx := q.db.Model(&entities.Product{}).Limit(1).Find(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &product, nil
}
