package queries

import "context"

type Tenant struct {
	ID                  string   `json:"id"`
	ProductID           string   `json:"product_id"`
	Product             *Product `json:"product,omitempty"`
	Name                string   `json:"name"`
	ResourceInformation string   `json:"resource_information"`
}

type TenantQuery interface {
	GetByID(ctx context.Context, id string) (Tenant, error)
	GetAllByOrganizationID(ctx context.Context, organizationID string) ([]Tenant, error)
}
