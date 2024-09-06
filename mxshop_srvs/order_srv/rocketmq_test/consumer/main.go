package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithGroupName("hb_test"),
	)

	if err := c.Subscribe("hb_mq_test", consumer.MessageSelector{}, func(c context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		//响应逻辑，类似于回调
		for _, msg := range msgs {
			fmt.Println("获取到数值是：%v\n", msg.String())
		}
		//直接返回success,此条消息被标记为已经消费过，不会再被消费
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println(err)
	}
	_ = c.Start()
	//不能让主携程退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
