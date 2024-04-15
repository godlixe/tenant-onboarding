package services

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"
)

type AuthService interface {
	Login(
		ctx context.Context,
		userRequest *entity.User,
	) (string, error)

	Register(
		ctx context.Context,
		user *entity.User,
	) (*entity.User, error)

	Me(
		ctx context.Context,
	) (*entity.User, error)
}
