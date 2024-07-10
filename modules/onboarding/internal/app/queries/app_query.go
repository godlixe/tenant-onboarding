package queries

import "context"

type App struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	FrontendURL string `json:"frontend_url"`
}

type AppQuery interface {
	GetAll(ctx context.Context) ([]App, error)
}
