package repositories

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
)

type BillingManagementRepository interface {
	CreateBilling(ctx context.Context, tenant *entities.Tenant) (string, error)
}
