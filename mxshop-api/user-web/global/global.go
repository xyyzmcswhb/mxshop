package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

//定义所有全局变量

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
	NacosConfig   *config.NacosConfig = &config.NacosConfig{}
)
