package postgresql

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/app/queries"

	"gorm.io/gorm"
)

type AppQuery struct {
	db *gorm.DB
}

func NewAppQuery(
	db *gorm.DB,
) *AppQuery {
	return &AppQuery{
		db: db,
	}
}

func (q *AppQuery) GetAll(ctx context.Context) ([]queries.App, error) {
	var apps []queries.App

	tx := q.db.Model(&queries.App{}).
		Find(&apps)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return apps, nil
}
