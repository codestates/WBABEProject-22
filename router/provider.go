package router

import (
	"github.com/gin-gonic/gin"

	"oos/controller"
	"oos/middleware"
)

func addProviderRoutes(rg *gin.RouterGroup) {
	provider := rg.Group("/provider")
	provider.Use(middleware.ValidateToken())
	provider.Use(middleware.ValidateScope("provider"))

	provider.POST("/products", controller.CreateProduct)
	provider.PUT("/products/:code", controller.UpdateProduct)
	provider.DELETE("/products/:code", controller.DeleteProduct)

	provider.GET("/orders", controller.ListOrders)
	provider.PUT("/orders/:id/status", controller.UpdateOrderStatus)

	provider.GET("/reviews/orders", controller.ListReviews)
}
