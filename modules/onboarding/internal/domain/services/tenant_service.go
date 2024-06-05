package services

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

type TenantService struct {
}

func NewTenantService() *TenantService {
	return &TenantService{}
}

func (s *TenantService) CreateNewTenant(
	ctx context.Context,
	name string,
	productID vo.ProductID,
	organizationID vo.OrganizationID,
) (*entities.Tenant, error) {
	return entities.NewTenant(
		vo.GenerateTenantID(),
		productID,
		organizationID,
		name,
		vo.TenantOnboarding,
	), nil
}
