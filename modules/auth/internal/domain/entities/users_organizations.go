package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UsersOrganizations struct {
	UserID         vo.UserID
	OrganizationID vo.OrganizationID
	Role           vo.OrganizationRole
}

func NewUsersOrganizations(
	userID vo.UserID,
	organizationID vo.OrganizationID,
	role vo.OrganizationRole,
) *UsersOrganizations {
	return &UsersOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
		Role:           role,
	}
}
