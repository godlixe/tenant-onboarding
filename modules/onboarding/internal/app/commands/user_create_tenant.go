package commands

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/framework"
)

type UserCreateTenantRequest struct {
	productID      string
	organizationID string
	name           string
}

func NewUserCreateTenantRequest(
	productID string,
	organizationID string,
	name string,
) UserCreateTenantRequest {
	return UserCreateTenantRequest{
		productID:      productID,
		organizationID: organizationID,
		name:           name,
	}
}

type UserCreateTenantCommand struct {
	infrastructureRepository   repositories.InfrastructureRepository
	tenantRepository           repositories.TenantRepository
	productRepository          repositories.ProductRepository
	tenantDeploymentRepository repositories.TenantDeploymentRepository
	tenantManagementRepository repositories.TenantManagementRepository
}

func NewUserCreateTenantCommand(
	infrastructureRepository repositories.InfrastructureRepository,
	tenantRepository repositories.TenantRepository,
	productRepository repositories.ProductRepository,
	tenantDeploymentRepository repositories.TenantDeploymentRepository,
	tenantManagementRepository repositories.TenantManagementRepository,
) *UserCreateTenantCommand {
	return &UserCreateTenantCommand{
		infrastructureRepository:   infrastructureRepository,
		tenantRepository:           tenantRepository,
		productRepository:          productRepository,
		tenantDeploymentRepository: tenantDeploymentRepository,
		tenantManagementRepository: tenantManagementRepository,
	}
}

func (c *UserCreateTenantCommand) Execute(ctx context.Context, r UserCreateTenantRequest) (*entities.Tenant, error) {
	productID, err := valueobjects.NewProductID(r.productID)
	if err != nil {
		return nil, err
	}

	organizationID, err := valueobjects.NewOrganizationID(r.organizationID)
	if err != nil {
		return nil, err
	}

	tenant := entities.NewTenant(
		valueobjects.GenerateTenantID(),
		productID,
		organizationID,
		r.name,
		valueobjects.TenantCreated,
	)

	// If app isusing integrated mode, call tenant management
	// API to store tenant data
	if framework.CheckIntegratedMode() {
		tenant, err = c.tenantManagementRepository.CreateTenant(ctx, tenant)
		if err != nil {
			return nil, err
		}
	} else {
		err = c.tenantRepository.Create(ctx, tenant)
		if err != nil {
			return nil, err
		}

		tenantDeploymentJob := entities.NewTenantDeploymentJob(
			tenant,
		)

		// Publishes data to standalone queue.
		err = c.tenantDeploymentRepository.PublishTenantDeploymentJob(ctx, tenantDeploymentJob)
		if err != nil {
			return nil, err
		}
	}

	return tenant, nil
}
