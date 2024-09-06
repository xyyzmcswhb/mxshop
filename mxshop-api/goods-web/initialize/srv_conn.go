package initialize

import (
	"fmt"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"

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
		zap.S().Fatal("【initsrvconn】连接用户服务失败")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}
