package queries

import "context"

type Organization struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type OrganizationQuery interface {
	GetAllUserOrganization(ctx context.Context, userID string) ([]Organization, error)
	GetOrganizationLevel(ctx context.Context, organizationID string, userID string) (string, error)
}
