package services

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/internal/domain/users/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	user *entity.User,
) (*entity.User, error) {
	err := s.userRepository.CreateUser(ctx, user)

	return user, err
}
