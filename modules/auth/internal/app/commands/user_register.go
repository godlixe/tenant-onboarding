package commands

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	"tenant-onboarding/modules/auth/internal/domain/errors"
	"tenant-onboarding/modules/auth/internal/domain/repositories"
	"tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UserRegisterRequest struct {
	name     string
	email    string
	username string
	password string
}

func NewUserRegisterRequest(
	name string,
	email string,
	username string,
	password string,
) UserRegisterRequest {
	return UserRegisterRequest{
		name:     name,
		email:    email,
		username: username,
		password: password,
	}
}

type UserRegisterCommand struct {
	userRepository repositories.UserRepository
}

func NewUserRegisterCommand(
	userRepository repositories.UserRepository,
) *UserRegisterCommand {
	return &UserRegisterCommand{
		userRepository: userRepository,
	}
}

func (c *UserRegisterCommand) Execute(ctx context.Context, r UserRegisterRequest) error {
	existingUser, err := c.userRepository.GetByEmail(ctx, r.email)
	if err != nil {
		return err
	}

	if (*existingUser != entities.User{}) {
		return errors.ErrEmailExists
	}

	user := entities.NewUser(
		valueobjects.GenerateUserID(),
		r.name,
		r.email,
		r.username,
		r.password,
	)

	err = user.HashPassword()
	if err != nil {
		return err
	}

	err = c.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
