package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "oos/docs"
	"oos/logger"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		auth := c.GetHeader("Authorization")
		// ADD AUTHENTICATION FUNCTION BELOW:
		fmt.Println("Authorization-word ", auth)
		c.Next() 
	}
}

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
	e.Use(CORS())
	e.Use(liteAuth())

	// Route groups
	v1 := e.Group("/v1")
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
