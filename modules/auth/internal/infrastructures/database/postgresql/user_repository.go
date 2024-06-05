package postgresql

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	userID vo.UserID,
) (*entities.User, error) {
	var user entities.User

	tx := r.db.Model(&entities.User{}).
		Where("id = ?", userID.String()).
		Limit(1).
		Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*entities.User, error) {
	var user entities.User

	tx := r.db.Model(&entities.User{}).
		Where("email = ?", email).
		Limit(1).
		Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *entities.User,
) error {
	tx := r.db.Model(&entities.User{}).
		Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
