package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/order_srv/global"
	"mxshop_srvs/order_srv/handler"
	"mxshop_srvs/order_srv/initialize"
	"mxshop_srvs/order_srv/proto"
	"mxshop_srvs/order_srv/utils"
	"mxshop_srvs/order_srv/utils/register/consul"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitRedis()
	initialize.InitSrvs()
	zap.S().Info(global.ServerConfig)

	IP := flag.String("ip", global.ServerConfig.Host, "server ip")
	Port := flag.Int("port", global.ServerConfig.Port, "server port")
	flag.Parse()
	if *Port == 0 {
		//若使用了默认值，则使用随机端口号覆盖
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Infof("server ip:%s port:%d", *IP, *Port)

	//fmt.Println("ip:", *IP, "port:", *Port)
	//启动grpc服务
	server := grpc.NewServer()
	proto.RegisterOrderServer(server, &handler.OrderServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检测
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//阻塞式方法，需改成异步调用,启动服务
	go func() {
		if err := server.Serve(lis); err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//服务注册
	srvRegisterClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := uuid.NewV4()
	serviceIdStr := fmt.Sprintf("%s", serviceId)
	err = srvRegisterClient.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceIdStr)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}
	zap.S().Infof("启动服务器，端口：%d", *Port) //拿到全局sugar,可以让我们自己设置一个全局的logger

	//监听订单超时topic
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("mxshop-order"),
	)

	if err := c.Subscribe("order_timeout", consumer.MessageSelector{}, handler.OrderTimeout); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出

	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_ = c.Shutdown()
	//注销服务
	if err = srvRegisterClient.DeRegister(serviceIdStr); err != nil {
		zap.S().Panic("服务注销失败", err.Error())
	} else {
		zap.S().Infof("服务注销成功")
	}
}
