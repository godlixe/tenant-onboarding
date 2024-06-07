package providers

import (
	"context"
	"tenant-onboarding/modules/onboarding/internal/app/commands"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/infrastructures/database/postgresql"
	"tenant-onboarding/modules/onboarding/internal/infrastructures/queue/pubsub"
	"tenant-onboarding/modules/onboarding/internal/presentation/controllers"
	"tenant-onboarding/modules/onboarding/internal/presentation/routes"
	"tenant-onboarding/modules/onboarding/internal/processor"
	"tenant-onboarding/providers"

	"github.com/samber/do"
)

func RegisterDependencies(app *providers.App) {
	productQuery := postgresql.NewProductQuery(app.DB)
	appQuery := postgresql.NewAppQuery(app.DB)

	infrastructureRepository := postgresql.NewInfrastructureRepository(app.DB)
	tenantInfrastructureRepository := postgresql.NewTenantsInfrastructuresRepository(app.DB)
	tenantRepository := postgresql.NewTenantRepository(app.DB)
	productRepository := postgresql.NewProductRepository(app.DB)
	tenantDeploymentRepository := pubsub.NewTenantDeploymentRepository(app.Queue)

	userCreateTenantCmd := commands.NewUserCreateTenantCommand(
		infrastructureRepository,
		tenantRepository,
		productRepository,
		tenantDeploymentRepository,
	)

	productController := controllers.NewProductController(productQuery)
	appController := controllers.NewAppController(appQuery)
	tenantController := controllers.NewTenantController(userCreateTenantCmd)

	routes.ProductRoutes(app.Webserver, productController)
	routes.AppRoutes(app.Webserver, appController)
	routes.TenantRoutes(app.Webserver, tenantController)

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.InfrastructureRepository, error) {
		return infrastructureRepository, nil
	})

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.TenantsInfrastructuresRepository, error) {
		return tenantInfrastructureRepository, nil
	})

	processor.Run(context.TODO(), app)
}
