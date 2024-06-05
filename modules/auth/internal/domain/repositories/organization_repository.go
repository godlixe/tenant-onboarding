package repositories

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type OrganizationRepository interface {
	GetByID(context.Context, vo.OrganizationID) (*entities.Organization, error)
	GetBySubdomain(context.Context, string) (*entities.Organization, error)
	Create(context.Context, *entities.Organization) error
}
