package postgresql

import (
	"context"
	"tenant-onboarding/modules/auth/internal/app/queries"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(
	db *gorm.DB,
) *UserQuery {
	return &UserQuery{
		db: db,
	}
}

func (q *UserQuery) GetByID(ctx context.Context, userID string) (*queries.User, error) {
	var user queries.User

	tx := q.db.Table("users").Where("id = ?", userID).Take(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
