package postgresql

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(
	db *gorm.DB,
) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

func (r *OrganizationRepository) GetByID(
	ctx context.Context,
	organizationID vo.OrganizationID,
) (*entities.Organization, error) {
	var organization entities.Organization

	tx := r.db.Model(&entities.Organization{}).
		Where("id = ?", organizationID.String()).
		Limit(1).
		Find(&organization)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organization, nil
}

func (r *OrganizationRepository) GetByIdentifier(
	ctx context.Context,
	identifier string,
) (*entities.Organization, error) {
	var organization entities.Organization

	tx := r.db.Model(&entities.Organization{}).
		Where("identifier = ?", identifier).
		Limit(1).
		Find(&organization)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organization, nil
}

func (r *OrganizationRepository) Create(
	ctx context.Context,
	organization *entities.Organization,
) error {
	tx := r.db.Model(&entities.Organization{}).
		Create(organization)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
