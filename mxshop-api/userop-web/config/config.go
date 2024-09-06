package config

type UserOpSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name         string          `mapstructure:"name" json:"name"`
	Host         string          `mapstructure:"host" json:"host"`
	Tags         []string        `mapstructure:"tags" json:"tags"`
	Port         int             `mapstructure:"port" json:"port"`
	JWTInfo      JWTConfig       `mapstructure:"jwt" json:"jwt"`
	GoodsSrvInfo UserOpSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	UserOpInfo   UserOpSrvConfig `mapstructure:"userop_srv" json:"userop_srv"`
	RedisInfo    RedisConfig     `mapstructure:"redis" json:"redis"`
	ConsulInfo   ConsulConfig    `mapstructure:"consul" json:"consul"`
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
