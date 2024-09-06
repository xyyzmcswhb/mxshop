package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT,PATCH") //放开所有http请求方法的权限
		//允许加的header
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//针对option方法做优化
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) //返回204
		}
	}
}