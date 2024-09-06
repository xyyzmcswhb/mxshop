package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/userop_srv/global"
	"mxshop_srvs/userop_srv/handler"
	"mxshop_srvs/userop_srv/initialize"
	"mxshop_srvs/userop_srv/proto"
	"mxshop_srvs/userop_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "server ip")
	Port := flag.Int("port", 50059, "server port")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	if *Port == 0 {
		//若使用了默认值，则使用随机端口号覆盖
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Infof("server ip:%s port:%d", *IP, *Port)

	//fmt.Println("ip:", *IP, "port:", *Port)
	//启动grpc服务
	server := grpc.NewServer()
	proto.RegisterAddressServer(server, &handler.UserOpServer{})
	proto.RegisterMessageServer(server, &handler.UserOpServer{})
	proto.RegisterUserFavServer(server, &handler.UserOpServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检测
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "127.0.0.1", *Port),
		Timeout:                        "5s",
		Interval:                       "5s",  //5s检查一次
		DeregisterCriticalServiceAfter: "15s", //10s有效期
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = []string{global.ServerConfig.Name}
	registration.Address = "127.0.0.1"
	registration.Check = check
	//1.如何启动两个服务 2.即使能够通过终端启动两个服务，但是注册时会被覆盖
	err = client.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}

	//阻塞式方法，需改成异步调用
	go func() {
		if err := server.Serve(lis); err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//注销服务
	if err = client.Agent().ServiceDeregister(registration.ID); err != nil {
		zap.S().Info("注销失败:", err.Error())
	}
	zap.S().Info("注销成功", zap.String("serviceID", registration.ID))
}
