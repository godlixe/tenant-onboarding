package providers

import (
	"context"
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/app/commands"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	tenantmanagement "tenant-onboarding/modules/onboarding/internal/infrastructures/api/tenant_management"
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
	tenantQuery := postgresql.NewTenantQuery(app.DB)

	infrastructureRepository := postgresql.NewInfrastructureRepository(app.DB)
	tenantInfrastructureRepository := postgresql.NewTenantsInfrastructuresRepository(app.DB)
	tenantRepository := postgresql.NewTenantRepository(app.DB)
	productRepository := postgresql.NewProductRepository(app.DB)
	tenantDeploymentRepository := pubsub.NewTenantDeploymentRepository(app.Queue)
	tenantOnboardedRepository := pubsub.NewTenantOnboardedRepository(app.Queue)
	tenantManagementRepository := tenantmanagement.NewTenantManagementRepository(http.DefaultClient)

	userCreateTenantCmd := commands.NewUserCreateTenantCommand(
		infrastructureRepository,
		tenantRepository,
		productRepository,
		tenantDeploymentRepository,
		tenantManagementRepository,
	)

	productController := controllers.NewProductController(productQuery)
	appController := controllers.NewAppController(appQuery)
	tenantController := controllers.NewTenantController(userCreateTenantCmd, tenantQuery)

	routes.ProductRoutes(app.Webserver, productController)
	routes.AppRoutes(app.Webserver, appController)
	routes.TenantRoutes(app.Webserver, tenantController)

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.InfrastructureRepository, error) {
		return infrastructureRepository, nil
	})

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.TenantOnboardedRepository, error) {
		return tenantOnboardedRepository, nil
	})

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.TenantsInfrastructuresRepository, error) {
		return tenantInfrastructureRepository, nil
	})

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.ProductRepository, error) {
		return productRepository, nil
	})

	do.Provide(app.Injector, func(injector *do.Injector) (repositories.TenantRepository, error) {
		return tenantRepository, nil
	})

	processor.Run(context.TODO(), app)
}
