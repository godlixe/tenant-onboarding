package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
)

type Organization struct {
	ID        vo.OrganizationID
	Name      string
	Subdomain string

	database.Timestamp
}

func NewOrganization(
	id vo.OrganizationID,
	name string,
	subdomain string,
) *Organization {
	return &Organization{
		ID:        id,
		Name:      name,
		Subdomain: subdomain,
	}
}
