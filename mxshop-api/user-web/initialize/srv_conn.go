package initialize

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	userConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=hb", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Fatal("【initsrvconn】连接用户服务失败")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
func InitSrvConn2() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	usersrvHost := ""
	usersrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "user-srv"`))
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	fmt.Println(global.ServerConfig.UserSrvInfo.Name)
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		usersrvHost = value.Address
		usersrvPort = value.Port
		break
		//fmt.Println(key)
	}
	if usersrvHost == "" {
		zap.S().Fatal("【initsrvconn】连接用户服务失败")
	}
	fmt.Println(usersrvHost, usersrvPort)
	//拨号连接用户grpc服务
	userConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", usersrvHost, usersrvPort),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), //轮询
	)
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】",
			"msg", err.Error(),
		)
	}
	//潜在问题：1.后续服务下线 2.端口号或者IP变更 （负载均衡）
	//3. 一个连接多个goroutine共用，是否会存在性能问题？（解决方案，连接池）
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
