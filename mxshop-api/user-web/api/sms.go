package api

import (
	"context"
	"fmt"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/rand"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// 短信验证码
func GenerateSmsCode(width int) string {
	//生成指定长度的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(uint64(time.Now().UnixNano()))

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)]) //写入
		//sb.WriteByte(numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	//var loginForm LoginForm
	if err := c.ShouldBind(&sendSmsForm); err != nil { //参数绑定
		HandleValidatorErrors(c, err)
		return
	}
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecret)
	if err != nil {
		panic(err)
	}
	smsCode := GenerateSmsCode(6)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile
	request.QueryParams["SignName"] = "暮雪生鲜"
	request.QueryParams["TemplateCode"] = "SMS_301275378"
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}"
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("response is %#v\n", response)
	//后面注册的时候需要验证码校验

	//保存验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second) //5分钟过期
	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}
