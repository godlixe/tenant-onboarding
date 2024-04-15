package services

import (
	"context"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/internal/domain/users/repository"
)

type TenantService struct {
	tenantRepository repository.TenantRepository
	tenantPublisher  repository.TenantPublisher
}

func NewTenantService(
	tenantRepository repository.TenantRepository,
	tenantPublisher repository.TenantPublisher,
) *TenantService {
	return &TenantService{
		tenantRepository: tenantRepository,
		tenantPublisher:  tenantPublisher,
	}
}

func (s *TenantService) CreateTenant(
	ctx context.Context,
	tenant *entity.Tenant,
) (*entity.Tenant, error) {
	// create tenant row in db
	err := s.tenantRepository.CreateTenant(ctx, tenant)
	if err != nil {
		return nil, err
	}
	// create user tenant...
	// push to queue
	err = s.tenantPublisher.Publish(ctx, tenant)
	if err != nil {
		return nil, err
	}

	return tenant, err
}
