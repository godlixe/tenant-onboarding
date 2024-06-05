package queries

import "context"

type Infrastructure struct {
	ID                 string
	ProductID          string
	Name               string
	InfrastructureType string
	UserCount          int
	UserLimit          int
	Metadata           string
}

type InfrastructureQuery interface {
	GetByProductIDInfraTypeOrdered(ctx context.Context, productID string) ([]Infrastructure, error)
}
