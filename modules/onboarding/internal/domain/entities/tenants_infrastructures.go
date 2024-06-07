package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
)

type TenantsInfrastructures struct {
	TenantID         vo.TenantID
	InfrastructureID vo.InfrastructureID

	database.Timestamp
}

func NewTenantsInfrastructures(
	tenantID vo.TenantID,
	infrastructureID vo.InfrastructureID,
) *TenantsInfrastructures {
	return &TenantsInfrastructures{
		TenantID:         tenantID,
		InfrastructureID: infrastructureID,
	}
}
