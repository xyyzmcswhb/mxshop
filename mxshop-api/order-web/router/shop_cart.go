package router

import (
	"mxshop-api/order-web/api/shopcart"
	"mxshop-api/order-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	//路由组
	//shopCartRouter := Router.Group("shopcarts")
	shopCartRouter := Router.Group("shopcarts").Use(middlewares.JWTAuth())
	{
		shopCartRouter.GET("list", shopcart.List)            //购物车列表
		shopCartRouter.POST("add", shopcart.New)             //添加条目
		shopCartRouter.DELETE("delete/:id", shopcart.Delete) //添加条目
		shopCartRouter.PATCH("update/:id", shopcart.Update)  //修改条目，全部修改用put，部分修改用patch

	}
	//{
	//	GoodsRouter.GET("list", middlewares.JWTAuth(), middlewares.IsadminAuth(), order.list)          //订单列表
	//	GoodsRouter.POST("add_new_order", middlewares.JWTAuth(), middlewares.IsadminAuth(), order.New) //此接口需要管理员权限 ,新建订单
	//	//GoodsRouter.POST("add_new_goods", goods.NewGoods) //此接口需要管理员权限
	//	GoodsRouter.GET("/:id/detail", middlewares.JWTAuth(), middlewares.IsadminAuth(), order.Detail) //订单详情
	//	GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth())                   //删除订单
	//}
}
