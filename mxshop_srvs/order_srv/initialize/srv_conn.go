package initialize

import (
	"fmt"
	"mxshop_srvs/order_srv/global"
	"mxshop_srvs/order_srv/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvs() {
	//初始化第三方微服务的连接client
	goodsConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Fatal("【initsrvconn】连接商品服务失败")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	//初始化库存服务
	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.InvSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)

	if err != nil {
		zap.S().Fatal("【initsrvconn】连接库存服务失败")
	}

	global.InventorySrvClient = proto.NewInventoryClient(inventoryConn)
}
