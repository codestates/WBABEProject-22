package router

import (
	"github.com/gin-gonic/gin"

	"oos/controller"
)

func addProviderRoutes(rg *gin.RouterGroup) {
	provider := rg.Group("/provider")
	
	provider.POST("/products", controller.CreateProduct)
	provider.PUT("/products/:code", controller.UpdateProduct)
	provider.DELETE("/products/:code", controller.DeleteProduct)
	
	provider.GET("/orders", controller.GetOrders)
	provider.PUT("/orders/:id/status", controller.UpdateOrderStatus)
	
	provider.GET("/reviews/orders", controller.GetReviews)
}
