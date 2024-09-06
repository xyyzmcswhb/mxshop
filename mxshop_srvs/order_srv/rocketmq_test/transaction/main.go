package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type OrderListener struct{}

func (o *OrderListener) ExecuteLocalTransaction(*primitive.Message) primitive.LocalTransactionState {
	//执行本地事务
	fmt.Println("execute localTransaction")
	time.Sleep(3 * time.Second)
	fmt.Println("execute localTransaction success")
	return primitive.CommitMessageState
}

func (o *OrderListener) CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState {
	//回查逻辑
	return primitive.RollbackMessageState
}

func main() {
	//事务
	p, err := rocketmq.NewTransactionProducer(
		&OrderListener{},
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	transaction, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("transformer", []byte("this is transaction message")))
	if err != nil {
		fmt.Printf("发送失败: %s", err.Error())
	} else {
		fmt.Println("send transaction success:%s\n", transaction.String())
	}

	time.Sleep(1 * time.Hour)
	if err := p.Shutdown(); err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
