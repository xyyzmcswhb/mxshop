package router

import (
	"mxshop-api/goods-web/api/banners"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	//路由组
	BannersRouter := Router.Group("banners")
	{
		BannersRouter.GET("list", banners.List)                                                             //轮播图列表
		BannersRouter.POST("add_new_banner", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.New) //此接口需要管理员权限，增加轮播图
		BannersRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.Delete)      //删除轮播图
		BannersRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.Update)         //更新轮播图
	}
}
