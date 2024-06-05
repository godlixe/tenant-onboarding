package queries

import "context"

type App struct {
	ID   string
	Name string
}

type AppQuery interface {
	GetAll(ctx context.Context) ([]App, error)
}
