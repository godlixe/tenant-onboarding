package postgresql

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindById(
	ctx context.Context,
	id uuid.UUID,
) (*entity.User, error) {
	var user entity.User

	tx := r.db.Model(entity.User{}).Where("id = ?", id).First(&user)

	return &user, tx.Error
}

func (r *UserRepository) FindByUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {
	var user entity.User

	tx := r.db.Model(entity.User{}).Where("username = ?", username).First(&user)

	return &user, tx.Error
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	var user entity.User

	tx := r.db.Model(entity.User{}).Where("email = ?", email).First(&user)

	return &user, tx.Error
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *entity.User,
) error {
	tx := r.db.Model(entity.User{}).Create(&user)

	return tx.Error
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	user *entity.User,
) error {
	tx := r.db.Model(entity.User{}).Updates(&user)

	return tx.Error
}
