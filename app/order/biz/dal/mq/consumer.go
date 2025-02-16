package mq

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Consumer struct {
	client   *RabbitClient
	prefetch int
}

func NewConsumer(client *RabbitClient, prefetch int) *Consumer {
	return &Consumer{
		client:   client,
		prefetch: prefetch,
	}
}

func StartConsumer(ctx context.Context, client *RabbitClient, prefetch int) {
	consumer := NewConsumer(client, prefetch)
	if err := consumer.ConsumePreOrders(ctx, handlePreOrder); err != nil {
		log.Printf("消费预占消息失败: %v", err)
	}
	if err := consumer.ConsumeOrders(ctx, handleOrder); err != nil {
		log.Printf("消费订单消息失败: %v", err)
	}
	if err := consumer.ConsumeDelay(ctx, handleDelayOrder); err != nil {
		log.Printf("消费延迟消息失败: %v", err)
	}
}

func (c *Consumer) ConsumePreOrders(ctx context.Context, handler func(context.Context, PreOrderMessage) error) error {
	return c.consume(ctx, PreOrderQueue, func(d amqp.Delivery) error {
		var msg PreOrderMessage
		if err := sonic.Unmarshal(d.Body, &msg); err != nil {
			return fmt.Errorf("消息解析失败: %w", err)
		}

		if err := handler(ctx, msg); err != nil {
			// 业务处理失败，重新入队
			d.Nack(false, true)
			return err
		}
		d.Ack(false)
		return nil
	})
}

func (c *Consumer) ConsumeOrders(ctx context.Context, handler func(context.Context, OrderMessage) error) error {
	return c.consume(ctx, OrderQueue, func(d amqp.Delivery) error {
		var msg OrderMessage
		if err := sonic.Unmarshal(d.Body, &msg); err != nil {
			d.Nack(false, false)
			return fmt.Errorf("订单消息解析失败: %w", err)
		}

		if err := handler(ctx, msg); err != nil {
			//if errors.Is(err, ErrTransient) {
			//	d.Nack(false, true)
			//} else {
			//	d.Nack(false, false)
			//}
			// todo 数据库级错误需要重试，业务逻辑错误不重试
			d.Nack(false, false)
			return err
		}
		d.Ack(false)
		return nil
	})
}

func (c *Consumer) ConsumeDelay(ctx context.Context, handler func(context.Context, DelayMessage) error) error {
	return c.consume(ctx, DelayQueue, func(d amqp.Delivery) error {
		var msg DelayMessage
		if err := sonic.Unmarshal(d.Body, &msg); err != nil {
			return fmt.Errorf("延迟消息解析失败: %w", err)
		}

		if err := handler(ctx, msg); err != nil {
			log.Printf("处理延迟消息失败: %v", err)
			d.Nack(false, false)
			return err
		}
		d.Ack(false)
		return nil
	})
}

func (c *Consumer) consume(ctx context.Context, queue string, handler func(amqp.Delivery) error) error {
	ch, err := c.client.Channel()
	if err != nil {
		return err
	}

	if err := ch.Qos(c.prefetch, 0, false); err != nil {
		return fmt.Errorf("设置QoS失败: %w", err)
	}

	msgs, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("启动消费者失败: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-msgs:
				if err := handler(d); err != nil {
					log.Printf("消息处理失败: %v", err)
				}
			}
		}
	}()

	return nil
}
