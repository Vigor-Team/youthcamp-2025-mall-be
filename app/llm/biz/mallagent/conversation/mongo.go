package conversation

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"

	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	"github.com/cloudwego/eino/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "conversations"
)

type MongoMemory struct {
	mu            sync.RWMutex
	client        *mongo.Client
	db            *mongo.Database
	collection    *mongo.Collection
	maxWindowSize int
	conversations sync.Map
	ctx           context.Context
}

type MongoMemoryConfig struct {
	Client        *mongo.Client
	DB            *mongo.Database
	Collection    *mongo.Collection
	MaxWindowSize int
	Ctx           context.Context
}

type Conversation struct {
	mu            sync.Mutex
	ID            string
	Messages      []*schema.Message
	maxWindowSize int
	collection    *mongo.Collection
}

func NewMongoMemory(cfg MongoMemoryConfig) *MongoMemory {
	if cfg.MaxWindowSize <= 0 {
		cfg.MaxWindowSize = 6
	}
	if cfg.Collection == nil {
		cfg.Collection = cfg.DB.Collection(collectionName)
	}
	return &MongoMemory{
		client:        cfg.Client,
		db:            cfg.DB,
		collection:    cfg.Collection,
		maxWindowSize: cfg.MaxWindowSize,
		ctx:           cfg.Ctx,
	}
}

func GetDefaultMemory(ctx context.Context) *MongoMemory {
	return NewMongoMemory(MongoMemoryConfig{
		Client:        dalmongo.Client,
		DB:            dalmongo.DB,
		MaxWindowSize: 6,
		Ctx:           ctx,
	})
}

func (m *MongoMemory) GetConversation(id string, createIfNotExist bool) *Conversation {
	if conv, ok := m.conversations.Load(id); ok {
		return conv.(*Conversation)
	}

	filter := bson.M{"_id": id}
	var convDoc struct {
		ID       string            `bson:"_id"`
		Messages []*schema.Message `bson:"messages"`
	}

	err := m.collection.FindOne(m.ctx, filter).Decode(&convDoc)
	switch {
	case err == nil:
		conv := &Conversation{
			ID:            convDoc.ID,
			Messages:      convDoc.Messages,
			maxWindowSize: m.maxWindowSize,
			collection:    m.collection,
		}
		m.conversations.Store(id, conv)
		return conv

	case err == mongo.ErrNoDocuments && createIfNotExist:
		newConv := &Conversation{
			ID:            id,
			Messages:      make([]*schema.Message, 0),
			maxWindowSize: m.maxWindowSize,
			collection:    m.collection,
		}

		updateOpts := options.Update().SetUpsert(true)
		update := bson.M{
			"$setOnInsert": bson.M{
				"messages": []*schema.Message{},
			},
		}
		if _, err := m.collection.UpdateOne(m.ctx, filter, update, updateOpts); err != nil {
			return nil
		}

		m.conversations.Store(id, newConv)
		return newConv

	default:
		klog.CtxErrorf(m.ctx, "failed to get conversation: %v", err)
		return nil
	}
}

func (m *MongoMemory) ListConversations() ([]string, error) {
	cursor, err := m.collection.Find(m.ctx, bson.D{},
		options.Find().SetProjection(bson.M{"_id": 1}))
	if err != nil {
		return nil, fmt.Errorf("failed to list conversations: %w", err)
	}
	defer cursor.Close(m.ctx)

	var ids []string
	for cursor.Next(m.ctx) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			continue
		}
		ids = append(ids, result.ID)
	}
	return ids, nil
}

func (m *MongoMemory) DeleteConversation(id string) error {
	if _, err := m.collection.DeleteOne(m.ctx, bson.M{"_id": id}); err != nil {
		return fmt.Errorf("failed to delete conversation: %w", err)
	}
	m.conversations.Delete(id)
	return nil
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
		return fmt.Errorf("failed to append message: %w", err)
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
