package main

import (
	"context"
	"fmt"
	"mxshop-api/user-web/global"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, 6379),
	})
	fmt.Printf("Redis Host: %s, Port: %d\n", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port)
	value, err := rdb.Get(context.Background(), "15206190234").Result()
	if err == redis.Nil {
		fmt.Println("key不存在")
	}
	fmt.Println(value)
}
