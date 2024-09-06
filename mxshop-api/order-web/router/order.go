package router

import (
	"mxshop-api/order-web/api/order"
	"mxshop-api/order-web/api/pay"
	"mxshop-api/order-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	//路由组
	OrderRouter := Router.Group("order").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("list", order.List)                                 //订单列表
		OrderRouter.POST("add_new_order", middlewares.JWTAuth(), order.New) //此接口需要管理员权限 ,新建订单
		OrderRouter.GET("/:id/detail", middlewares.JWTAuth(), order.Detail) //订单详情
		//BannersRouter.POST("add_new_banner", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.New) //此接口需要管理员权限，增加轮播图
		//BannersRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.Delete)      //删除轮播图
		//BannersRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), banners.Update)         //更新轮播图
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}

}
