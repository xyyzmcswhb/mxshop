package initialize

import (
	"fmt"
	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {

	goodsConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Fatal("【initsrvconn】连接商品服务失败")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	//userOpConn, err := grpc.NewClient(
	//	fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserOpInfo.Name),
	//	grpc.WithInsecure(),
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	//)

	userOpConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserOpInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户操作服务失败】")
	}

	global.UserFavClient = proto.NewUserFavClient(userOpConn)
	global.MessageClient = proto.NewMessageClient(userOpConn)
	global.AddressClient = proto.NewAddressClient(userOpConn)
}
