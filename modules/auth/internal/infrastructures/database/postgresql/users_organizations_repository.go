package postgresql

import (
	"context"
	"tenant-onboarding/modules/auth/internal/domain/entities"
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"

	"gorm.io/gorm"
)

type UsersOrganizations struct {
	db *gorm.DB
}

func NewUsersOrganizationsRepository(
	db *gorm.DB,
) *UsersOrganizations {
	return &UsersOrganizations{
		db: db,
	}
}

func (r *UsersOrganizations) GetByUserIDOrgID(
	ctx context.Context,
	userID vo.UserID,
	organizationID vo.OrganizationID,
) (*entities.UsersOrganizations, error) {
	var userOrg entities.UsersOrganizations
	tx := r.db.Model(&entities.UsersOrganizations{}).
		Where("user_id = ?", userID).
		Where("organization_id = ?", organizationID).
		First(&userOrg)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &userOrg, nil
}

func (r *UsersOrganizations) Create(
	ctx context.Context,
	userOrg *entities.UsersOrganizations,
) error {
	tx := r.db.Model(&entities.UsersOrganizations{}).
		Create(&userOrg)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
