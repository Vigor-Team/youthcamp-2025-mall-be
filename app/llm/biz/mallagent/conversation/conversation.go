package conversation

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/cloudwego/eino/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Conversation struct {
	mu            sync.Mutex
	ID            string
	UserID        string
	Messages      []*schema.Message
	maxWindowSize int
	collection    *mongo.Collection
}

type ConversationDoc struct {
	ID       string            `bson:"_id"`
	UserID   string            `bson:"userId"`
	Messages []*schema.Message `bson:"messages"`
	Created  time.Time         `bson:"createdAt"`
	Updated  time.Time         `bson:"updatedAt"`
}

func (c *Conversation) Append(msg *schema.Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	update := bson.M{
		"$push": bson.M{
			"messages": bson.M{
				"$each":     []*schema.Message{msg},
				"$slice":    -c.maxWindowSize,
				"$position": 0,
			},
		},
	}
	if _, err := c.collection.UpdateByID(
		context.Background(),
		c.ID,
		update,
		options.Update().SetUpsert(true),
	); err != nil {
		return consts.ErrAppendMessage
	}

	c.Messages = append([]*schema.Message{msg}, c.Messages...)
	if len(c.Messages) > c.maxWindowSize {
		c.Messages = c.Messages[:c.maxWindowSize]
	}
	return nil
}

func (c *Conversation) GetMessages() []*schema.Message {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.Messages) <= c.maxWindowSize {
		return c.Messages
	}
	return c.Messages[:c.maxWindowSize]
}

func (c *Conversation) GetFullMessages() []*schema.Message {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Messages
}
