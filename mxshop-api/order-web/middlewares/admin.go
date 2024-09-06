package middlewares

import (
	"mxshop-api/order-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsadminAuth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			//非管理员权限
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next()
	})
}
