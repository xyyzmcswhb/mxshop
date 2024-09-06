package initialize

import (
	"fmt"
	"mxshop_srvs/order_srv/global"

	goredislib "github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis"
)

func InitRedis() {
	global.Client = goredislib.NewClient(&goredislib.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	pool := goredis.NewPool(global.Client)
	global.Rs = redsync.New(pool)
}
