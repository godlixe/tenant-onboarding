package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
)

type Organization struct {
	ID         vo.OrganizationID
	Name       string
	Identifier string

	database.Timestamp
}

func NewOrganization(
	id vo.OrganizationID,
	name string,
	identifier string,
) *Organization {
	return &Organization{
		ID:         id,
		Name:       name,
		Identifier: identifier,
	}
}
