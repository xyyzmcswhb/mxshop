package global

import (
	"mxshop-api/userop-web/config"
	"mxshop-api/userop-web/proto"

	ut "github.com/go-playground/universal-translator"
)

//定义所有全局变量

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	AddressClient  proto.AddressClient
	UserFavClient  proto.UserFavClient
	MessageClient  proto.MessageClient
	GoodsSrvClient proto.GoodsClient
	NacosConfig    *config.NacosConfig = &config.NacosConfig{}
)
