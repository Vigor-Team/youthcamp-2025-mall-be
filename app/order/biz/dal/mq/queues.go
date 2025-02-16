package mq

import (
	"fmt"
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
	// 死信交换机
	if err := ch.ExchangeDeclare(
		DLXExchange,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("声明死信交换机失败: %w", err)
	}

	// 主交换机
	if err := ch.ExchangeDeclare(
		MainExchange,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("声明主交换机失败: %w", err)
	}

	// queue
	// 预占队列
	if _, err := ch.QueueDeclare(
		PreOrderQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("声明预占队列失败: %w", err)
	}

	// 中间队列
	if _, err := ch.QueueDeclare(
		DelayQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    DLXExchange,
			"x-dead-letter-routing-key": "order_timeout",
			"x-message-ttl":             int32(10000), // 10s
		},
	); err != nil {
		return fmt.Errorf("声明中间队列失败: %w", err)
	}

	// 订单队列
	if _, err := ch.QueueDeclare(
		OrderQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("声明订单队列失败: %w", err)
	}

	// 死信队列
	if _, err := ch.QueueDeclare(
		DLXQueue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("声明延迟队列失败: %w", err)
	}

	// binding
	bindings := []struct {
		Queue    string
		Key      string
		Exchange string
	}{
		{PreOrderQueue, "pre_order", MainExchange},
		{OrderQueue, "create_order", MainExchange},
		{DelayQueue, "delay", MainExchange},
		{DLXQueue, "order_timeout", DLXExchange},
	}

	for _, b := range bindings {
		if err := ch.QueueBind(
			b.Queue,
			b.Key,
			b.Exchange,
			false,
			nil,
		); err != nil {
			return fmt.Errorf("绑定队列失败: %w", err)
		}
	}

	return nil
}
