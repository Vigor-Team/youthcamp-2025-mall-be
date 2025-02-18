package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	PreOrderQueue = "seckill.pre_order"
	OrderQueue    = "seckill.order_create"
	DelayQueue    = "seckill.delay"
	DLXQueue      = "seckill.dlx"
	DLXExchange   = "seckill.dlx_exchange"
	MainExchange  = "seckill.main_exchange"
)

type QueueManager struct {
	client *RabbitClient
}

func NewQueueManager(client *RabbitClient) *QueueManager {
	return &QueueManager{client: client}
}

func (qm *QueueManager) SetupQueues() error {
	ch, err := qm.client.Channel()
	if err != nil {
		return err
	}

	// exchange
	// dlx
	if err = ch.ExchangeDeclare(
		DLXExchange,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		klog.Errorf("failed to declare dlx exchange: %v", err)
	}

	// main
	if err = ch.ExchangeDeclare(
		MainExchange,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		klog.Errorf("failed to declare main exchange: %v", err)
		return err
	}

	// queue
	// pre order
	if _, err := ch.QueueDeclare(
		PreOrderQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		klog.Errorf("failed to declare pre order queue: %v", err)
		return err
	}

	// delay
	if _, err = ch.QueueDeclare(
		DelayQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    DLXExchange,
			"x-dead-letter-routing-key": "order_timeout",
			"x-message-ttl":             int32(20000),
		},
	); err != nil {
		klog.Errorf("failed to declare delay queue: %v", err)
		return err
	}

	// order
	if _, err = ch.QueueDeclare(
		OrderQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		klog.Errorf("failed to declare order queue: %v", err)
		return err
	}

	// dlx
	if _, err = ch.QueueDeclare(
		DLXQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		klog.Errorf("failed to declare dlx queue: %v", err)
		return err
	}

	// binding
	bindings := []struct {
		Queue    string
		Key      string
		Exchange string
	}{
		{PreOrderQueue, "pre_order", MainExchange},
		{OrderQueue, "order_create", MainExchange},
		{DelayQueue, "delay", MainExchange},
		{DLXQueue, "order_timeout", DLXExchange},
	}

	for _, b := range bindings {
		if err = ch.QueueBind(
			b.Queue,
			b.Key,
			b.Exchange,
			false,
			nil,
		); err != nil {
			klog.Errorf("failed to bind queue: %v", err)
			return err
		}
	}

	return nil
}
