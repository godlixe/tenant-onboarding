package repositories

import (
	"context"
	"tenant-onboarding/pkg/deployer/types"
)

type TenantOnboardedRepository interface {
	PublishTenantOnboarded(context.Context, *types.TenantOnboardedEvent) error
}
