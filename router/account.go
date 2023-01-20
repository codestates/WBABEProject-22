package router

import (
	"github.com/gin-gonic/gin"

	"oos/controller"
)

func addAccountRoutes(rg *gin.RouterGroup) {
	account := rg.Group("/account")
	account.POST("/login/:role", controller.Login)
}