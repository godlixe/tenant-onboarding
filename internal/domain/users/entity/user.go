package entity

import (
	"tenant-onboarding/pkg/auth"
	"tenant-onboarding/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	database.Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	hashedPassword, err := auth.HashAndSalt(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

func NewUser(
	Name string,
	Email string,
	Username string,
	Password string,
) *User {
	return &User{
		ID:       uuid.New(),
		Name:     Name,
		Email:    Email,
		Username: Username,
		Password: Password,
	}
}
