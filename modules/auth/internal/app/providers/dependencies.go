package providers

import (
	"tenant-onboarding/modules/auth/internal/app/commands"
	"tenant-onboarding/modules/auth/internal/infrastructures/database/postgresql"
	"tenant-onboarding/modules/auth/internal/presentation/controllers"
	"tenant-onboarding/modules/auth/internal/presentation/routes"
	"tenant-onboarding/providers"
)

func RegisterDependencies(app *providers.App) {
	db := app.DB
	engine := app.Webserver

	userQuery := postgresql.NewUserQuery(db)
	organizationQuery := postgresql.NewOrganizationQuery(db)

	userRepository := postgresql.NewUserRepository(db)
	organizationRepository := postgresql.NewOrganizationRepository(db)
	usersOrganizationsRepository := postgresql.NewUsersOrganizationsRepository(db)

	userLoginCmd := commands.NewUserLoginCommand(userRepository)
	userRegisterCmd := commands.NewUserRegisterCommand(userRepository)
	userCreateOrganizationCmd := commands.NewUserCreateOrganizationCommand(organizationRepository, usersOrganizationsRepository)

	authController := controllers.NewAuthController(userRegisterCmd, userLoginCmd, userQuery)
	organizationController := controllers.NewOrganizationController(userCreateOrganizationCmd, organizationQuery)

	routes.AuthRoutes(engine, authController)
	routes.OrganizationRoutes(engine, organizationController)
}
