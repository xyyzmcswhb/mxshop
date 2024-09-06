package router

import (
	"mxshop-api/goods-web/api/category"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	//路由组
	CategoryRouter := Router.Group("category")
	{
		CategoryRouter.GET("list", category.List)                                                               //分类列表
		CategoryRouter.POST("add_new_category", middlewares.JWTAuth(), middlewares.IsadminAuth(), category.New) //此接口需要管理员权限，增加分类
		CategoryRouter.GET("/:id", category.Detail)                                                             //分类详情
		CategoryRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), category.Delete)        //删除分类
		CategoryRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsadminAuth(), category.Update)           //更新分类
	}
}
