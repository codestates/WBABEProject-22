package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "oos/docs"
	"oos/logger"
	"oos/middleware"
)

func Engine() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	e := gin.Default()

	// Default middleware
	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())

	// Custom middleware
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(middleware.CORS())

	// Route groups
	v1 := e.Group("/v1")
	addAccountRoutes(v1)
	addCustomerRoutes(v1)
	addProviderRoutes(v1)

	// Swagger documentation
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("Start server")

	return e
}

// References
// Class material: lecture 12
// https://github.com/gin-gonic/examples/tree/master/group-routes
