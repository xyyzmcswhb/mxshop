package main

import (
	"context"
	"fmt"
	"mxshop_srvs/user_srv/proto"

	"google.golang.org/grpc"
)

var (
	userClient proto.UserClient
	conn       *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	userClient = proto.NewUserClient(conn)
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Nickname: fmt.Sprintf("huangbiao%d", i),
			Password: "generic password",
			Mobile:   fmt.Sprintf("1520619022%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserlist() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{ //客户端测试

		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.Nickname, user.Password)
		checkrsp, err := userClient.CheckUserPassword(context.Background(), &proto.PasswordCheck{
			Password:          "generic password",
			Encryptedpassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkrsp.Success)
	}
}

func main() {
	Init()
	//TestCreateUser()
	TestGetUserlist()
	conn.Close()

}
