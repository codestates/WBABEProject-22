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
	/* [코드리뷰]
	* 코드의 확장성을 고려하였을때, endpoint 관리가 잘 고려된 매우 좋은 코드라고 생각됩니다.
	* API 서버 내 코드 수정이 발생했을 경우, 리뉴얼 버전에 따라 
	* v01 방식의 클라이언트와, v02 방식의 클라이언트를 모두 받아줄 수 있는 
	* 확장성 있는 코드를 잘 만들어주셨습니다. 훌륭하다고 생각됩니다.
	*/
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
