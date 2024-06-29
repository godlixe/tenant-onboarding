package routes

import (
	"tenant-onboarding/middlewares"
	"tenant-onboarding/modules/auth/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(router *gin.Engine, organizationController *controllers.OrganizationController) {
	organizationRoutes := router.Group("/organization", middlewares.Authenticate())
	{
		organizationRoutes.GET("", organizationController.GetAllUserOrganization)
		organizationRoutes.POST("", organizationController.CreateOrganizations)
		organizationRoutes.GET("level", organizationController.GetOrganizationLevel)
	}
}
