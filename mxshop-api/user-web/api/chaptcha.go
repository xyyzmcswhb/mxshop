package api

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

//图形验证码

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80) //图形验证码属性
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s, //图片路径
	})

}
