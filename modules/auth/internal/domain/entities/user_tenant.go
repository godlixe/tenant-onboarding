package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UserTenant struct {
	UserID   vo.UserID
	TenantID vo.TenantID
	Role     string
}

func NewUserTenant(
	userID vo.UserID,
	tenantID vo.TenantID,
	role string,
) *UserTenant {
	return &UserTenant{
		UserID:   userID,
		TenantID: tenantID,
		Role:     role,
	}
}
