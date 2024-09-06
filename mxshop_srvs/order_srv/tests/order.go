package main

import (
	"context"
	"fmt"
	"mxshop_srvs/order_srv/proto"
)

func TestCreateCartItem(userid, goodid int32) {
	rsp, err := OrderClient.CreateCartItem(context.Background(), &proto.CartItemReq{
		UserId:  userid,
		GoodsId: goodid,
		Nums:    1,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("设置商品库存成功", rsp.Id)
}

func TestCartItemList(userId int32) {
	rsp, err := OrderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("拉取购物车列表成功")
	for _, item := range rsp.Data {
		fmt.Println(item)
	}
}

func TestUpdateCartStatus(Id int32) {
	_, err := OrderClient.UpdateCartItem(context.Background(), &proto.CartItemReq{
		Id:      Id,
		Checked: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("更新购物车商品状态成功")
}

func TestCreateOrder() {
	rsp, err := OrderClient.CreateOrder(context.Background(), &proto.OrderReq{
		UserId:  31,
		Address: "上海市",
		Name:    "咸鱼一族名存实亡",
		Mobile:  "15200000",
		Post:    "99",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Id)
}

func TestOrderDetail(orderId int32) {
	rsp, err := OrderClient.OrderDetail(context.Background(), &proto.OrderReq{
		Id: orderId,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
