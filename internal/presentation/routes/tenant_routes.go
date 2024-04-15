package routes

import (
	"tenant-onboarding/internal/presentation/controllers"
	"tenant-onboarding/internal/presentation/middlewares"

	"github.com/gin-gonic/gin"
)

func TenantRoutes(router *gin.Engine, tenantController *controllers.TenantController) {
	tenantRoutes := router.Group("/tenant", middlewares.Authenticate())
	{
		tenantRoutes.POST("create", tenantController.CreateTenant)
		// tenantRoutes.GET("me", middlewares.Authenticate(), authController.Me)
	}
}
