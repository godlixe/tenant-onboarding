package entities

import (
	vo "tenant-onboarding/modules/auth/internal/domain/valueobjects"
)

type UserTenant struct {
	UserID   vo.UserID
	TenantID vo.TenantID
	level    string
}

func NewUserTenant(
	userID vo.UserID,
	tenantID vo.TenantID,
	level string,
) *UserTenant {
	return &UserTenant{
		UserID:   userID,
		TenantID: tenantID,
		level:    level,
	}
}
