package mq

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Producer struct {
	client *RabbitClient
}

func NewProducer(client *RabbitClient) *Producer {
	return &Producer{client: client}
}

func (p *Producer) PublishPreOrder(ctx context.Context, msg PreOrderMessage) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		klog.CtxErrorf(ctx, "sonic.Marshal.err: %v", err)
		return err
	}

	return p.publish(ctx, MainExchange, "pre_order", body, nil)
}

func (p *Producer) PublishOrder(ctx context.Context, msg OrderMessage) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		klog.CtxErrorf(ctx, "sonic.Marshal.err: %v", err)
		return err
	}

	return p.publish(ctx, MainExchange, "order_create", body, nil)
}

func (p *Producer) PublishDelay(ctx context.Context, msg DelayMessage, delay time.Duration) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		klog.CtxErrorf(ctx, "sonic.Marshal.err: %v", err)
		return err
	}

	return p.publish(ctx, MainExchange, "delay", body, nil)
}

func (p *Producer) publish(ctx context.Context, exchange, key string, body []byte, headers amqp.Table) error {
	ch, err := p.client.Channel()
	if err != nil {
		klog.CtxErrorf(ctx, "client.Channel.err: %v", err)
		return err
	}

	return ch.PublishWithContext(ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Headers:      headers,
		},
	)
}
