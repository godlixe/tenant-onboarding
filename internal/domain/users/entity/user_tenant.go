package entity

import (
	"tenant-onboarding/pkg/database"

	"github.com/google/uuid"
)

type UserTenant struct {
	UserID   uuid.UUID `json:"user_id"`
	User     User      `json:"user"`
	TenantID uuid.UUID `json:"tenant_id"`
	Tenant   Tenant    `json:"tenant"`
	Role     string    `json:"role"`
	database.Timestamp
}
