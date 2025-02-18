package mq

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/consts"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
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
		klog.CtxErrorf(ctx, "consumer.ConsumePreOrders.err: %v", err)
	}
	if err := consumer.ConsumeOrders(ctx, handleOrder); err != nil {
		klog.CtxErrorf(ctx, "consumer.ConsumeOrders.err: %v", err)
	}
	if err := consumer.ConsumeDelay(ctx, handleDelayOrder); err != nil {
		klog.CtxErrorf(ctx, "consumer.ConsumeDelay.err: %v", err)
	}
}

func (c *Consumer) ConsumePreOrders(ctx context.Context, handler func(context.Context, PreOrderMessage) error) error {
	return c.consume(ctx, PreOrderQueue, func(d amqp.Delivery) error {
		var msg PreOrderMessage
		if err := sonic.Unmarshal(d.Body, &msg); err != nil {
			d.Nack(false, false)
			return consts.ErrConsumeMessage
		}

		if err := handler(ctx, msg); err != nil {
			klog.CtxErrorf(ctx, "ConsumePreOrders.handler.err: %v", err)
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
			return consts.ErrConsumeMessage
		}

		// todo retry
		if err := handler(ctx, msg); err != nil {
			klog.CtxErrorf(ctx, "ConsumeOrders.handler.err: %v", err)
			d.Nack(false, false)
			return err
		}
		d.Ack(false)
		return nil
	})
}

func (c *Consumer) ConsumeDelay(ctx context.Context, handler func(context.Context, DelayMessage) error) error {
	return c.consume(ctx, DLXQueue, func(d amqp.Delivery) error {
		var msg DelayMessage
		if err := sonic.Unmarshal(d.Body, &msg); err != nil {
			d.Nack(false, false)
			return consts.ErrConsumeMessage
		}

		if err := handler(ctx, msg); err != nil {
			klog.CtxErrorf(ctx, "ConsumeDelay.handler.err: %v", err)
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

	if err = ch.Qos(c.prefetch, 0, false); err != nil {
		klog.CtxErrorf(ctx, "ch.Qos.err: %v", err)
		return err
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
		klog.CtxErrorf(ctx, "ch.Consume.err: %v", err)
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-msgs:
				if err := handler(d); err != nil {
					klog.CtxErrorf(ctx, "consume.handler.err: %v", err)
				}
			}
		}
	}()

	return nil
}
