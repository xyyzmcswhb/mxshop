package main

import (
	"mxshop_srvs/userop_srv/proto"

	"google.golang.org/grpc"
)

var (
	userFavClient proto.UserFavClient
	addressClient proto.AddressClient
	messageClient proto.MessageClient
	conn          *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	userFavClient = proto.NewUserFavClient(conn)
	addressClient = proto.NewAddressClient(conn)
	messageClient = proto.NewMessageClient(conn)
}

func main() {
	Init()
	TestUserFavList()
	TestMessageList()
	TestAddressList()
	conn.Close()

}
