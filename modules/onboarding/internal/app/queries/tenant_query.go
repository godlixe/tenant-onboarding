package queries

import "context"

type Tenant struct {
	ID        string
	ProjectID string
	ProductID string
	Name      string
	Subdomain string
	JSON      string
}

type TenantQuery interface {
	GetByID(ctx context.Context, id string) (Tenant, error)
}
