package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"
)

func TestGetBrandList() {
	//for i := 0; i < 10; i++ {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}

	//}
	//rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{ //客户端测试
	//
	//	Pn:    1,
	//	PSize: 5,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for _, user := range rsp.Data {
	//	fmt.Println(user.Mobile, user.Nickname, user.Password)
	//	checkrsp, err := userClient.CheckUserPassword(context.Background(), &proto.PasswordCheck{
	//		Password:          "generic password",
	//		Encryptedpassword: user.Password,
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(checkrsp.Success)
	//}
}
