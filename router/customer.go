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
/* [코드리뷰]
* Group을 사용하여 API 성격에 따라 request를 관리하는 코드는 매우 좋은 코드입니다.
* 일반적으로 현업에서도 이와 같은 코드를 자주 사용합니다. 훌륭합니다.
*/
