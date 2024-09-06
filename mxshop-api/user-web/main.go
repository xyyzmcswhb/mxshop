package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	"mxshop-api/user-web/utils"
	"mxshop-api/user-web/utils/register/consul"
	"os"
	"os/signal"
	"syscall"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"

	"github.com/spf13/viper"

	ut "github.com/go-playground/universal-translator"

	myvalidator "mxshop-api/user-web/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

func main() {
	//port := 8081
	//1.初始化logger
	initialize.InitLogger()

	//2.初始化配置文件
	initialize.InitializeConfig()

	//3. 初始化router
	Router := initialize.Routers()

	logger, _ := zap.NewDevelopment()

	//4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5.初始化grpc服务连接
	initialize.InitSrvConn2()

	viper.AutomaticEnv()
	//如果是本地开发环境，端口号固定，否则随机获取端口号
	debug := viper.GetBool("MXSHOP_DEBUG")
	if debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//6.注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error { //解决mobile的翻译问题
			return ut.Add("mobile", "{0} 手机号码非法！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	zap.ReplaceGlobals(logger)
	srvRegisterClient := consul.NewRegisterClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId, _ := uuid.NewV4()
	serviceIdstr := fmt.Sprintf("%s", serviceId)
	err := srvRegisterClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceIdstr)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}
	/*
		s()可以获取一个全局的logger
		日志是分级别的，debug，info，warn,error,fetal
		s函数和L函数很有用，提供安全访问logger的途径，并发安全
	*/
	zap.S().Infof("启动服务器，端口：%d", global.ServerConfig.Port) //拿到全局sugar,可以让我们自己设置一个全局的logger

	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = srvRegisterClient.Deregister(serviceIdstr); err != nil {
		zap.S().Panic("服务注销失败", err.Error())
	} else {
		zap.S().Infof("服务注销成功")
	}
}
