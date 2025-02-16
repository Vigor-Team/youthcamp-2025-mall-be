package mq

import (
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	URL          string
	mu           sync.Mutex
	notifyClose  chan *amqp.Error
	shutdownChan chan struct{}
}

func NewRabbitClient(url string) *RabbitClient {
	return &RabbitClient{
		URL:          url,
		shutdownChan: make(chan struct{}),
	}
}

func (rc *RabbitClient) Connect() error {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	var err error
	rc.conn, err = amqp.Dial(rc.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	rc.channel, err = rc.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	rc.notifyClose = make(chan *amqp.Error)
	rc.channel.NotifyClose(rc.notifyClose)

	go rc.handleReconnect()
	return nil
}

func (rc *RabbitClient) handleReconnect() {
	for {
		select {
		case err := <-rc.notifyClose:
			if err != nil {
				log.Printf("连接异常关闭: %v，尝试重连...", err)
				_ = rc.Connect()
			}
		case <-rc.shutdownChan:
			return
		}
	}
}

func (rc *RabbitClient) Close() {
	close(rc.shutdownChan)
	_ = rc.channel.Close()
	_ = rc.conn.Close()
}

func (rc *RabbitClient) Channel() (*amqp.Channel, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if rc.channel == nil || rc.channel.IsClosed() {
		if err := rc.Connect(); err != nil {
			return nil, err
		}
	}
	return rc.channel, nil
}
