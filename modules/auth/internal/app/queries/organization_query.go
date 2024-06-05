package queries

import "context"

type Organization struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Subdomain string `json:"subdomain"`
}

type OrganizationQuery interface {
	GetAllUserOrganization(ctx context.Context, userID string) ([]Organization, error)
}
