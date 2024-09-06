package router

import (
	"mxshop-api/userop-web/api/userfav"
	"mxshop-api/userop-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfavs")
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuth(), userfav.Delete) // 删除收藏记录
		UserFavRouter.GET("/:id", middlewares.JWTAuth(), userfav.Detail)    // 获取收藏记录
		UserFavRouter.POST("", middlewares.JWTAuth(), userfav.New)          //新建收藏记录
		UserFavRouter.GET("", middlewares.JWTAuth(), userfav.List)          //获取当前用户的收藏
	}
}
