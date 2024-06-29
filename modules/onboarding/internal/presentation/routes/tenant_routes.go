package routes

import (
	"tenant-onboarding/middlewares"
	"tenant-onboarding/modules/onboarding/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func TenantRoutes(router *gin.Engine, tenantController *controllers.TenantController) {
	tenantRoutes := router.Group("/tenant", middlewares.Authenticate())
	{
		tenantRoutes.POST("create", tenantController.CreateTenant)
		tenantRoutes.GET("", middlewares.Authenticate(), tenantController.GetTenants)
	}
}
