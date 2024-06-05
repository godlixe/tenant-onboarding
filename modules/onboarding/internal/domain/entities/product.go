package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
)

type Product struct {
	ID               vo.ProductID
	AppID            int
	TierID           int
	DeploymentSchema string
	database.Timestamp
}
