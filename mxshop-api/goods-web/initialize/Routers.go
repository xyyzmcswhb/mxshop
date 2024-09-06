package initialize

import (
	"mxshop-api/goods-web/middlewares"
	router2 "mxshop-api/goods-web/router"

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
	ApiGroup := router.Group("/g/v1")
	router2.InitGoodsRouter(ApiGroup)
	router2.InitCategoryRouter(ApiGroup)
	router2.InitBannerRouter(ApiGroup)
	router2.InitBrandsRouter(ApiGroup)

	return router

}
