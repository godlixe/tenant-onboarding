package postgresql

import (
	"context"
	"tenant-onboarding/modules/auth/internal/app/queries"
	"tenant-onboarding/modules/auth/internal/domain/entities"

	"gorm.io/gorm"
)

type OrganizationQuery struct {
	db *gorm.DB
}

func NewOrganizationQuery(
	db *gorm.DB,
) *OrganizationQuery {
	return &OrganizationQuery{
		db: db,
	}
}

func (q *OrganizationQuery) GetAllUserOrganization(ctx context.Context, userID string) ([]queries.Organization, error) {
	var organizations []queries.Organization

	tx := q.db.Model(&entities.Organization{}).
		Joins("JOIN users_organizations ON organizations.id = users_organizations.organization_id").
		Where("users_organizations.user_id = ?", userID).
		Find(&organizations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return organizations, nil
}

func (q *OrganizationQuery) GetOrganizationLevel(ctx context.Context, organizationID string, userID string) (string, error) {
	var level string

	tx := q.db.Model(&entities.UsersOrganizations{}).
		Where("organization_id = ?", organizationID).
		Where("user_id = ?", userID).
		Select("level").Find(&level)

	if tx.Error != nil {
		return "", tx.Error
	}

	return level, nil
}
