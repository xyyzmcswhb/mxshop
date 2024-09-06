package config

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	Tags         []string       `mapstructure:"tags" json:"tags"`
	Port         int            `mapstructure:"port" json:"port"`
	OrderSrvInfo OrderSrvConfig `mapstructure:"order_srv" json:"order_srv"`
	GoodsSrvInfo GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	InvSrvInfo   GoodsSrvConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	JWTInfo      JWTConfig      `mapstructure:"jwt" json:"jwt"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	ConsulInfo   ConsulConfig   `mapstructure:"consul" json:"consul"`
	AlipayInfo   AlipayConfig   `mapstructure:"alipay" json:"alipay"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	goods     string `mapstructure:"goods"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"data_id"`
	Group     string `mapstructure:"group"`
}
