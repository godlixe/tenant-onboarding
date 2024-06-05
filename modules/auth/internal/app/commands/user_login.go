package commands

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/errors"
	"tenant-onboarding/modules/auth/internal/domain/repositories"
	"tenant-onboarding/pkg/auth"
)

type UserLoginRequest struct {
	email    string
	password string
}

func NewUserLoginRequest(
	email string,
	password string,
) UserLoginRequest {
	return UserLoginRequest{
		email:    email,
		password: password,
	}
}

type UserLoginCommand struct {
	userRepository repositories.UserRepository
}

func NewUserLoginCommand(
	userRepository repositories.UserRepository,
) *UserLoginCommand {
	return &UserLoginCommand{
		userRepository: userRepository,
	}
}

func (c *UserLoginCommand) Execute(ctx context.Context, r UserLoginRequest) (*string, error) {
	user, err := c.userRepository.GetByEmail(ctx, r.email)
	if err != nil {
		return nil, err
	}

	passwordMatch, err := auth.ComparePassword(user.Password, []byte(r.password))
	if err != nil {
		return nil, err
	}

	if !passwordMatch {
		return nil, errors.ErrEmailExists
	}

	token, err := auth.GenerateJWTToken(
		user.ID.String(),
	)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
