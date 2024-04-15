package repository

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindById(
		ctx context.Context,
		id uuid.UUID,
	) (*entity.User, error)

	FindByUsername(
		ctx context.Context,
		username string,
	) (*entity.User, error)

	CreateUser(
		ctx context.Context,
		user *entity.User,
	) error

	UpdateUser(
		ctx context.Context,
		user *entity.User,
	) error
}
