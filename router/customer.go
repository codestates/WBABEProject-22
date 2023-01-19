package router

import (
	"github.com/gin-gonic/gin"

	"oos/controller"
	"oos/middleware"
)

func addCustomerRoutes(rg *gin.RouterGroup) {
	customer := rg.Group("/customer")
	customer.Use(middleware.ValidateToken())
	customer.Use(middleware.ValidateScope("customer"))

	customer.GET("/products", controller.ListProducts)
	customer.GET("/products/:code", controller.GetProduct)

	customer.GET(":username/orders/active", controller.ListOrdersActive)
	customer.GET(":username/orders/history", controller.ListOrdersHistory)

	customer.GET("/orders/:id", controller.GetOrder)
	customer.POST("/orders", controller.CreateOrder)
	customer.PUT("/orders/:id/cart", controller.UpdateOrderItems)
	customer.DELETE("/orders/:id/cart", controller.DeleteOrderItems)
	customer.GET("/orders/:id/status", controller.GetOrderStatus)

	customer.POST("/reviews/orders/:id", controller.CreateReview)
	customer.GET("/reviews/products/:code", controller.ListReviewsProduct)
}
