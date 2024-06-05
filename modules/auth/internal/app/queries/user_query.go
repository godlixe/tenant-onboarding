package queries

import "context"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserQuery interface {
	GetByID(ctx context.Context, userID string) (*User, error)
}
