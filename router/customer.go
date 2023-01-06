package router

import (
	"github.com/gin-gonic/gin"

	"oos/controller"
)

func addCustomerRoutes(rg *gin.RouterGroup) {
	customer := rg.Group("/customer")
	
	customer.GET("/products", controller.GetProducts)
	customer.GET("/products/:code", controller.GetProduct)
	
	customer.GET(":username/orders/active", controller.GetOrdersActive)
	customer.GET(":username/orders/history", controller.GetOrdersHistory)

	customer.GET("/orders/:id", controller.GetOrder)
	customer.POST("/orders", controller.CreateOrder)
	customer.PUT("/orders/:id/cart", controller.UpdateOrderItems)
	customer.DELETE("/orders/:id/cart", controller.DeleteOrderItems)
	customer.GET("/orders/:id/status", controller.GetOrderStatus)
	
	customer.POST("/reviews/orders/:id", controller.CreateReview)
	customer.GET("/reviews/products/:code", controller.GetReview)
}
