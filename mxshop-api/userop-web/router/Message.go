package router

import (
	"mxshop-api/userop-web/api/message"
	"mxshop-api/userop-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("list", message.List)    // 获取所有的留言
		MessageRouter.POST("add_new", message.New) //新建轮播图
	}
}
