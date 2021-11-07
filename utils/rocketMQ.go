package utils

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"red_envelope/configure"
)

var MyProducer rocketmq.Producer
var MyConsumer rocketmq.PushConsumer

//init函数会自动执行
func init() {
	//初始化异步生产者
	var err error
	MyProducer, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{configure.RMQ_NAMESEVER_ADDR})),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewManualQueueSelector()))
	if err != nil {
		if configure.GIN_MODE == "debug" {
			fmt.Printf("create producer error: %s\n", err.Error())
		}
		MyLog.Error("create producer error: ", err.Error())
		os.Exit(1)
	}
	err = MyProducer.Start()
	if err != nil {
		if configure.GIN_MODE == "debug" {
			fmt.Printf("start producer error: %s\n", err.Error())
		}
		MyLog.Error("start producer error: ", err.Error())
		os.Exit(1)
	}

	//初始化topic
	rmqWalletTopic := "rmq_wallet"
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{configure.RMQ_NAMESEVER_ADDR})))

	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(rmqWalletTopic),
		admin.WithBrokerAddrCreate(configure.RMQ_BROKER_ADDR),
	)
	if err != nil {
		if configure.GIN_MODE == "debug" {
			fmt.Printf("create topic error: %s\n", err.Error())
		}
		MyLog.Error("create topic error: ", err.Error())
		os.Exit(1)
	}

	//初始化消费者
	MyConsumer, _ := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{configure.RMQ_NAMESEVER_ADDR})),
	)
	err = MyConsumer.Subscribe(rmqWalletTopic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	err = MyConsumer.Start()
	if err != nil {
		if configure.GIN_MODE == "debug" {
			fmt.Printf("start consumer error: %s\n", err.Error())
		}
		MyLog.Error("start consumer error: ", err.Error())
		os.Exit(1)
	}
}

type rocketMqMessage struct {
	messageId    int64
	messageBytes string
	tag          string
	topic        string
}

func SendToRMQ(msg rocketMqMessage) {
	err := MyProducer.SendAsync(context.Background(),
		func(ctx context.Context, result *primitive.SendResult, e error) {
			if e != nil {
				if configure.GIN_MODE == "debug" {
					fmt.Printf("receive message error: %s\n", e)
				}
				MyLog.Error("receive message error: ", e)
			}
		}, primitive.NewMessage(msg.topic, []byte(msg.messageBytes)))
	if err != nil {
		if configure.GIN_MODE == "debug" {
			fmt.Printf("send message error: %s\n", err)
		}
		MyLog.Error("send message error: ", err)
	}
}

func RecvFromRMQ() []byte {
	return nil
}
