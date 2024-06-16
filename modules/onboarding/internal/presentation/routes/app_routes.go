package routes

import (
	"tenant-onboarding/modules/onboarding/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, appController *controllers.AppController) {
	appRoutes := router.Group("/app")
	{
		appRoutes.GET("", appController.GetAll)
	}
}
