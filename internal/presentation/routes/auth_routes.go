package routes

import (
	"tenant-onboarding/internal/presentation/controllers"
	"tenant-onboarding/internal/presentation/middlewares"

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
