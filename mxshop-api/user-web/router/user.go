package router

import (
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	//路由组
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsadminAuth(), api.GetUserList)
		UserRouter.POST("password_login", api.PasswordLogin)
		UserRouter.POST("Register", api.RegisterUser)
	}

	//服务注册和发现

}
