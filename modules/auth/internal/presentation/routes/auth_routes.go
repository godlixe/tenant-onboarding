package routes

import (
	"tenant-onboarding/middlewares"
	"tenant-onboarding/modules/auth/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("register", authController.Register)
		authRoutes.POST("login", authController.Login)
		authRoutes.GET("me", middlewares.Authenticate(), authController.Me)
	}
}
