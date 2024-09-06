package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"

	"github.com/golang/protobuf/ptypes/empty"
)

func TestGetAllCategorysList() {
	//for i := 0; i < 10; i++ {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.JsonData)
}

func TestGetSubCategorysList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 130358,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.SubCategorys)
}
