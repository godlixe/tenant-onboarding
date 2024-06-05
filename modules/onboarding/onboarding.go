package onboarding

import (
	"tenant-onboarding/modules/onboarding/internal/app/providers"
	appProvider "tenant-onboarding/providers"
)

func RegisterModule(app *appProvider.App) {
	providers.RegisterDependencies(app)
}
