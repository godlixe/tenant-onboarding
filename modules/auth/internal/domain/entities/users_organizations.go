package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UsersOrganizations struct {
	UserID         vo.UserID
	OrganizationID vo.OrganizationID
	Level          vo.OrganizationLevel
}

func NewUsersOrganizations(
	userID vo.UserID,
	organizationID vo.OrganizationID,
	level vo.OrganizationLevel,
) *UsersOrganizations {
	return &UsersOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
		Level:          level,
	}
}
