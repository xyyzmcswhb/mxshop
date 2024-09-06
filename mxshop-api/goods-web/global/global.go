package global

import (
	"mxshop-api/goods-web/config"
	"mxshop-api/goods-web/proto"

	ut "github.com/go-playground/universal-translator"
)

//定义所有全局变量

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
	NacosConfig    *config.NacosConfig = &config.NacosConfig{}
)
