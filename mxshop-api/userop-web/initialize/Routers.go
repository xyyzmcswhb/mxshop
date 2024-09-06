package initialize

import (
	"mxshop-api/userop-web/middlewares"
	router2 "mxshop-api/userop-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	})
	//配置跨域
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/uop/v1")
	router2.InitAddressRouter(ApiGroup)
	router2.InitUserFavRouter(ApiGroup)
	router2.InitMessageRouter(ApiGroup)

	return router

}
