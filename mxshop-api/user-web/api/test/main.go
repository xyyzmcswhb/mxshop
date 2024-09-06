package test

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "*", "*")
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = "15206190327"
	request.QueryParams["SignName"] = "暮雪生鲜"
	request.QueryParams["TemplateCode"] = "SMS_301275378"
	request.QueryParams["TemplateParam"] = "{\"code\":" + "777777" + "}"
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("response is %#v\n", response)
}
