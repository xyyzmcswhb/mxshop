package main

import (
	"mxshop_srvs/inventory_srv/proto"
	"sync"

	"google.golang.org/grpc"
)

var (
	InvClient proto.InventoryClient
	conn      *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	InvClient = proto.NewInventoryClient(conn)
}

func main() {
	Init()
	//var i int32
	//for i = 421; i <= 840; i++ {
	//	TestSetInv(i, 100)
	//}
	//TestCreateUser()
	//TestGetBrandList()
	//TestGetSubCategorysList()
	//TestGetCategoryBrandList()
	//TestGetAllCategorysList()
	//TestGetGoodsList()
	//var i int32
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go TestSell(&wg)
	}
	wg.Wait()
	//TestInvDetail(421)
	//TestSell()
	//TestReback()
	err := conn.Close()
	if err != nil {
		return
	}

}
