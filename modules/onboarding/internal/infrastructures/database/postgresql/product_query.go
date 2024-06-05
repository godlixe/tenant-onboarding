package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/app/queries"

	"gorm.io/gorm"
)

type ProductQuery struct {
	db *gorm.DB
}

func NewProductQuery(
	db *gorm.DB,
) *ProductQuery {
	return &ProductQuery{
		db: db,
	}
}

func (q *ProductQuery) GetProducts(ctx context.Context, filter *queries.ProductFilter) ([]queries.Product, error) {
	var products []queries.Product

	tx := q.db.Model(&queries.Product{}).Preload("App").Preload("Tier")

	if filter.AppID > 0 {
		tx.Where("app_id = ?", filter.AppID)
	}

	tx.Debug().Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}

func (q *ProductQuery) GetProductsByAppID(ctx context.Context, appID int) ([]queries.Product, error) {
	var products []queries.Product

	tx := q.db.Model(&queries.Product{}).
		Where("app_id = ?", appID).
		Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}
