package auth

import (
	"tenant-onboarding/modules/auth/internal/app/providers"
	appProvider "tenant-onboarding/providers"
)

func RegisterModule(app *appProvider.App) {
	providers.RegisterDependencies(app)
}
