package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/inventory_srv/global"
	"mxshop_srvs/inventory_srv/handler"
	"mxshop_srvs/inventory_srv/initialize"
	"mxshop_srvs/inventory_srv/proto"
	"mxshop_srvs/inventory_srv/utils"
	"mxshop_srvs/inventory_srv/utils/register/consul"
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
	proto.RegisterInventoryServer(server, &handler.InventoryServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检测
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	srvRegisterClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := uuid.NewV4()
	serviceIdstr := fmt.Sprintf("%s", serviceId)
	err = srvRegisterClient.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceIdstr)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}
	zap.S().Infof("启动服务器，端口：%d", *Port) //拿到全局sugar,可以让我们自己设置一个全局的logger

	//阻塞式方法，需改成异步调用,启动服务
	go func() {
		if err := server.Serve(lis); err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//监听库存归还topic
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("mxshop-inventory-mq"),
	)

	//消息订阅
	if err := c.Subscribe("inventory_return", consumer.MessageSelector{}, handler.AutoInvReturn); err != nil {
		fmt.Println(err)
	}
	_ = c.Start()
	//不能让主携程退出
	//time.Sleep(time.Hour)
	//_ = c.Shutdown()
	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_ = c.Shutdown()
	//注销服务
	if err = srvRegisterClient.DeRegister(serviceIdstr); err != nil {
		zap.S().Panic("服务注销失败", err.Error())
	} else {
		zap.S().Infof("服务注销成功")
	}
}
