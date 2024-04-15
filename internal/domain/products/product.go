package products

import (
	"tenant-onboarding/pkg/database"

	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID `json:"id"`
	AppID  int       `json:"app_id"`
	TierID int       `json:"tier_id"`
	DeploymentSchema
	database.Timestamp
}
