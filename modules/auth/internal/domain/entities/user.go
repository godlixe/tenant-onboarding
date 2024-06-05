package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
	"tenant-onboarding/pkg/auth"
	"tenant-onboarding/pkg/database"
)

type User struct {
	ID       vo.UserID `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Email    string
	Username string
	Password string

	database.Timestamp
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	u.ID = vo.GenerateUserID()

// 	hashedPassword, err := auth.HashAndSalt(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	u.Password = hashedPassword

// 	return nil
// }

func NewUser(
	id vo.UserID,
	name string,
	email string,
	username string,
	password string,
) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := auth.HashAndSalt(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}
