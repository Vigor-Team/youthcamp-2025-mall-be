package mq

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/conf"
	"os"
)

var Client *RabbitClient

func Init() {
	url := fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), conf.GetConf().MQ.Address)
	Client = NewRabbitClient(url)

	if err := Client.Connect(); err != nil {
		panic(err)
	}

	qm := NewQueueManager(Client)
	if err := qm.SetupQueues(); err != nil {
		panic(err)
	}

	StartConsumer(context.Background(), Client, consts.MqConsumerPreFetchCount)
}
