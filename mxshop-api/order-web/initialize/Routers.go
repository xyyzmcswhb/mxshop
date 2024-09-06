package initialize

import (
	"mxshop-api/order-web/middlewares"
	router2 "mxshop-api/order-web/router"

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
	ApiGroup := router.Group("/o/v1")
	router2.InitOrderRouter(ApiGroup)
	router2.InitShopCartRouter(ApiGroup)
	//router2.InitGoodsRouter(ApiGroup)
	//router2.InitCategoryRouter(ApiGroup)
	//router2.InitBannerRouter(ApiGroup)
	//router2.InitBrandsRouter(ApiGroup)

	return router

}
