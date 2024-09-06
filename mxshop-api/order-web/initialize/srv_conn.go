package initialize

import (
	"fmt"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"

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

	OrderConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Fatal("【initsrvconn】连接订单服务失败")
	}

	global.OrderSrvClient = proto.NewOrderClient(OrderConn)

	InvConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.InvSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Fatal("【initsrvconn】连接库存服务失败")
	}

	global.InvSrvClient = proto.NewInventoryClient(InvConn)
}
