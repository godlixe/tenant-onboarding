package routes

import (
	"tenant-onboarding/modules/onboarding/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, productController *controllers.ProductController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", productController.GetAll)
		// tenantRoutes.GET("me", middlewares.Authenticate(), authController.Me)
	}
}
