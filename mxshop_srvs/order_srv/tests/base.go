package main

import (
	"mxshop_srvs/order_srv/proto"

	"google.golang.org/grpc"
)

var (
	OrderClient proto.OrderClient
	conn        *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	OrderClient = proto.NewOrderClient(conn)
}

func main() {
	Init()
	//TestCreateCartItem(31, 421)
	//TestInvDetail(421)
	TestCartItemList(31)
	//TestUpdateCartStatus(1)
	//TestSell()
	//TestCreateOrder()
	//TestOrderDetail(5)
	//TestReback()
	err := conn.Close()
	if err != nil {
		return
	}
}
