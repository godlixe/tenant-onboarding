package commands

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
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
}

func NewUserCreateTenantCommand(
	infrastructureRepository repositories.InfrastructureRepository,
	tenantRepository repositories.TenantRepository,
	productRepository repositories.ProductRepository,
	tenantDeploymentRepository repositories.TenantDeploymentRepository,
) *UserCreateTenantCommand {
	return &UserCreateTenantCommand{
		infrastructureRepository:   infrastructureRepository,
		tenantRepository:           tenantRepository,
		productRepository:          productRepository,
		tenantDeploymentRepository: tenantDeploymentRepository,
	}
}

func (c *UserCreateTenantCommand) Execute(ctx context.Context, r UserCreateTenantRequest) error {
	productID, err := valueobjects.NewProductID(r.productID)
	if err != nil {
		return err
	}

	organizationID, err := valueobjects.NewOrganizationID(r.organizationID)
	if err != nil {
		return err
	}

	tenant := entities.NewTenant(
		valueobjects.GenerateTenantID(),
		productID,
		organizationID,
		r.name,
		valueobjects.TenantOnboarding,
	)

	err = c.tenantRepository.Create(ctx, tenant)
	if err != nil {
		return err
	}
	product, err := c.productRepository.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	tenantDeploymentJob := entities.NewTenantDeploymentJob(
		tenant,
		product,
	)

	err = c.tenantDeploymentRepository.PublishTenantDeploymentJob(ctx, tenantDeploymentJob)
	if err != nil {
		return err
	}

	return nil
}
