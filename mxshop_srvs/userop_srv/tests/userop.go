package main

import (
	"context"
	"fmt"
	"mxshop_srvs/userop_srv/proto"
)

func TestAddressList() {
	rsp, err := addressClient.GetAddressList(context.Background(), &proto.AddressRequest{
		UserId: 31,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestMessageList() {
	rsp, err := messageClient.MessageList(context.Background(), &proto.MessageRequest{
		UserId: 31,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestUserFavList() {
	rsp, err := userFavClient.GetFavList(context.Background(), &proto.UserFavRequest{
		UserId: 31,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
