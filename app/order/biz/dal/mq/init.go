package mq

import "context"

var Client *RabbitClient

func Init() {
	Client = NewRabbitClient("amqp://guest:guest@localhost:5672/")

	if err := Client.Connect(); err != nil {
		panic(err)
	}

	qm := NewQueueManager(Client)
	if err := qm.SetupQueues(); err != nil {
		panic(err)
	}

	StartConsumer(context.Background(), Client, 10)
}
