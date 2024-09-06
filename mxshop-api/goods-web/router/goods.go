package router

import (
	"mxshop-api/goods-web/api/goods"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	//路由组
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("list", goods.List)                                                                 //商品列表
		GoodsRouter.POST("add_new_goods", middlewares.JWTAuth(), middlewares.IsadminAuth(), goods.NewGoods) //此接口需要管理员权限
		//GoodsRouter.POST("add_new_goods", goods.NewGoods) //此接口需要管理员权限
		GoodsRouter.GET("/:id", goods.Detail)                                                           //商品详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), goods.DeleteGoods) //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                    //获取商品库存信息
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), goods.UpdateStatus) //更新商品状态
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), goods.UpdateGoods)    //更新商品信息
	}
}
