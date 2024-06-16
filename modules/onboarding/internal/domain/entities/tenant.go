package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
	"tenant-onboarding/pkg/events/domain"
)

type Tenant struct {
	ID               vo.TenantID
	ProductID        vo.ProductID
	OrganizationID   vo.OrganizationID
	Name             string
	Status           vo.TenantStatus
	InfrastructureID vo.InfrastructureID

	events []domain.Event
	database.Timestamp
}

func NewTenant(
	id vo.TenantID,
	productID vo.ProductID,
	organizationID vo.OrganizationID,
	name string,
	status vo.TenantStatus,
) *Tenant {
	return &Tenant{
		ID:             id,
		ProductID:      productID,
		OrganizationID: organizationID,
		Name:           name,
		Status:         status,
	}
}
