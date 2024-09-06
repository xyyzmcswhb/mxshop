package global

import (
	"mxshop-api/order-web/config"
	"mxshop-api/order-web/proto"

	ut "github.com/go-playground/universal-translator"
)

//定义所有全局变量

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	OrderSrvClient proto.OrderClient
	GoodsSrvClient proto.GoodsClient
	InvSrvClient   proto.InventoryClient
	NacosConfig    *config.NacosConfig = &config.NacosConfig{}
)
