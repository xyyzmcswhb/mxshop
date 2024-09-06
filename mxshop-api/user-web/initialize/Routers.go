package initialize

import (
	"mxshop-api/user-web/middlewares"
	router2 "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	//router.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"code":    200,
	//		"success": true,
	//	})
	//})
	//配置跨域
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup) //初始化图形验证码信息
	return router

}
