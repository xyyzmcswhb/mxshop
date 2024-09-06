package main

import (
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

var (
	brandClient proto.GoodsClient
	conn        *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	//TestCreateUser()
	//TestGetBrandList()
	//TestGetSubCategorysList()
	//TestGetCategoryBrandList()
	//TestGetAllCategorysList()
	//TestGetGoodsList()
	//TestBatchGetGoods()
	TestGetGoodsDetail()
	err := conn.Close()
	if err != nil {
		return
	}

}
