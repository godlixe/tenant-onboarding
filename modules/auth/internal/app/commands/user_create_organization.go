package commands

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	"tenant-onboarding/modules/auth/internal/domain/errors"
	"tenant-onboarding/modules/auth/internal/domain/repositories"
	"tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UserCreateOrganizationRequest struct {
	name      string
	subdomain string
	userID    string
}

func NewUserCreateOrganizationRequest(
	name string,
	subdomain string,
	userID string,
) UserCreateOrganizationRequest {
	return UserCreateOrganizationRequest{
		name:      name,
		subdomain: subdomain,
		userID:    userID,
	}
}

type UserCreateOrganizationCommand struct {
	organizationRepository       repositories.OrganizationRepository
	usersOrganizationsRepository repositories.UsersOrganizationsRepository
}

func NewUserCreateOrganizationCommand(
	organizationRepository repositories.OrganizationRepository,
	usersOrganizationsRepository repositories.UsersOrganizationsRepository,
) *UserCreateOrganizationCommand {
	return &UserCreateOrganizationCommand{
		organizationRepository:       organizationRepository,
		usersOrganizationsRepository: usersOrganizationsRepository,
	}
}

func (c *UserCreateOrganizationCommand) Execute(ctx context.Context, r UserCreateOrganizationRequest) error {
	existingOrg, err := c.organizationRepository.GetBySubdomain(ctx, r.subdomain)
	if err != nil {
		return err
	}
	if (*existingOrg != entities.Organization{}) {
		return errors.ErrOrgSubdomainExists
	}

	organization := entities.NewOrganization(
		valueobjects.GenerateOrganizationID(),
		r.name,
		r.subdomain,
	)

	err = c.organizationRepository.Create(ctx, organization)
	if err != nil {
		return err
	}

	userID, err := valueobjects.NewUserID(r.userID)
	if err != nil {
		return err
	}

	usersOrganization := entities.NewUsersOrganizations(
		userID,
		organization.ID,
		valueobjects.RoleOwner,
	)

	err = c.usersOrganizationsRepository.Create(ctx, usersOrganization)
	if err != nil {
		return err
	}

	return nil
}
