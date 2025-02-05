package conversation

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const (
	collectionName = "conversations"
)

type Bucket struct {
	mu            sync.RWMutex
	client        *mongo.Client
	db            *mongo.Database
	collection    *mongo.Collection
	maxWindowSize int
	conversations sync.Map
	ctx           context.Context
}

type BucketConfig struct {
	Client        *mongo.Client
	DB            *mongo.Database
	Collection    *mongo.Collection
	MaxWindowSize int
	Ctx           context.Context
}

func NewBucket(cfg BucketConfig) *Bucket {
	if cfg.MaxWindowSize <= 0 {
		cfg.MaxWindowSize = 6
	}
	if cfg.Collection == nil {
		cfg.Collection = cfg.DB.Collection(collectionName)
	}
	return &Bucket{
		client:        cfg.Client,
		db:            cfg.DB,
		collection:    cfg.Collection,
		maxWindowSize: cfg.MaxWindowSize,
		ctx:           cfg.Ctx,
	}
}

func GetDefaultBucket(ctx context.Context) *Bucket {
	return NewBucket(BucketConfig{
		Client:        dalmongo.Client,
		DB:            dalmongo.DB,
		MaxWindowSize: 6,
		Ctx:           ctx,
	})
}

func (m *Bucket) GetConversation(conversationId, userId string) (*Conversation, error) {
	if cached, ok := m.conversations.Load(conversationId); ok {
		conv := cached.(*Conversation)
		if conv.UserID != userId {
			klog.CtxErrorf(m.ctx, "conversation id %s already exists with different user id %s", conversationId, conv.UserID)
			return nil, consts.ErrGetConversation
		}
		return conv, nil
	}

	filter := bson.M{
		"_id":    conversationId,
		"userId": userId,
	}

	var doc ConversationDoc
	err := m.collection.FindOne(m.ctx, filter).Decode(&doc)
	switch {
	case err == nil:
		conv := &Conversation{
			ID:            doc.ID,
			UserID:        userId,
			Messages:      doc.Messages,
			maxWindowSize: m.maxWindowSize,
			collection:    m.collection,
		}
		m.conversations.Store(conversationId, conv)
		return conv, nil

	case err == mongo.ErrNoDocuments:
		newConv := &Conversation{
			ID:            conversationId,
			UserID:        userId,
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
			return nil, consts.ErrGetConversation
		}

		m.conversations.Store(conversationId, newConv)
		return newConv, nil

	default:
		klog.CtxErrorf(m.ctx, "failed to get conversation: %v", err)
		return nil, consts.ErrGetConversation
	}
}

func (m *Bucket) ListConversations(userId string) ([]string, error) {
	cursor, err := m.collection.Find(m.ctx, bson.M{"userId": userId}, options.Find().SetProjection(bson.M{"_id": 1}))
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

func (m *Bucket) DeleteConversation(conversationId, userId string) error {
	filter := bson.M{
		"_id":    conversationId,
		"userId": userId,
	}

	res, err := m.collection.DeleteOne(m.ctx, filter)
	if err != nil {
		klog.CtxErrorf(m.ctx, "failed to delete conversation: %v", err)
		return consts.ErrDeleteConversation
	}

	if res.DeletedCount == 0 {
		klog.CtxErrorf(m.ctx, "attempted to delete non-existent conversation %s", conversationId)
		return consts.ErrDeleteConversation
	}

	m.conversations.Delete(conversationId)
	return nil
}
