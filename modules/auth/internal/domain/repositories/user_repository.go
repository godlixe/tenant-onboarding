package repositories

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UserRepository interface {
	GetByID(context.Context, vo.UserID) (*entities.User, error)
	GetByEmail(context.Context, string) (*entities.User, error)
	Create(context.Context, *entities.User) error
}
