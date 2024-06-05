package repositories

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UsersOrganizationsRepository interface {
	Create(context.Context, *entities.UsersOrganizations) error
	GetByUserIDOrgID(context.Context, vo.UserID, vo.OrganizationID) (*entities.UsersOrganizations, error)
}
