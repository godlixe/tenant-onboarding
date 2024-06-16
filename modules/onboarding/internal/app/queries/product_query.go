package queries

import "context"

type Tier struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Product struct {
	ID       string `json:"id"`
	AppID    int    `json:"app_id"`
	App      *App   `json:"app,omitempty"`
	TierName string `json:"tier_name"`
	Price    int    `json:"price"`
}

type ProductFilter struct {
	AppID int `form:"app_id"`
}

type ProductQuery interface {
	GetProducts(ctx context.Context, filter *ProductFilter) ([]Product, error)
}
