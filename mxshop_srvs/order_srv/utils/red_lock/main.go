package main

import (
	"fmt"
	"time"

	"sync"

	goredislib "github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis"
)

var (
	client *goredislib.Client
	rs     *redsync.Redsync
)

func init() {
	client = goredislib.NewClient(&goredislib.Options{
		Addr: "127.0.0.1:6379",
	})
	pool := goredis.NewPool(client)
	rs = redsync.New(pool)
}

func main() {
	gNum := 2
	mutexname := "421"

	var wg sync.WaitGroup
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func(id int) {
			defer wg.Done()
			mutex := rs.NewMutex(mutexname)

			for {
				fmt.Printf("Goroutine %d: 开始获取锁\n", id)
				if err := mutex.Lock(); err != nil {
					fmt.Printf("Goroutine %d: 获取锁失败, 错误: %v\n", id, err)
					time.Sleep(time.Second) // 重试之前先等待一段时间
					continue
				}

				fmt.Printf("Goroutine %d: 获取锁成功\n", id)
				time.Sleep(time.Second * 5)

				fmt.Printf("Goroutine %d: 开始释放锁\n", id)
				if ok, err := mutex.Unlock(); !ok || err != nil {
					fmt.Printf("Goroutine %d: 释放锁失败, 错误: %v\n", id, err)
				} else {
					fmt.Printf("Goroutine %d: 释放锁成功\n", id)
				}
				break
			}
		}(i)
	}
	wg.Wait()
}
