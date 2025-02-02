package mem

//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//
//	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
//	"github.com/cloudwego/eino/schema"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//const (
//	collectionName   = "conversations"
//	defaultBatchSize = 100
//)
//
//type MongoMemory struct {
//	mu            sync.RWMutex
//	client        *mongo.Client
//	db            *mongo.Database
//	collection    *mongo.Collection
//	maxWindowSize int
//	conversations sync.Map
//	ctx           context.Context
//}
//
//type MongoMemoryConfig struct {
//	Client        *mongo.Client
//	DB            *mongo.Database
//	MaxWindowSize int
//	Ctx           context.Context
//}
//
//type Conversation struct {
//	ID            string            `bson:"_id"`
//	Messages      []*schema.Message `bson:"messages"`
//	maxWindowSize int
//	collection    *mongo.Collection
//	mu            sync.Mutex
//}
//
//func NewMongoMemory(cfg MongoMemoryConfig) *MongoMemory {
//	if cfg.MaxWindowSize <= 0 {
//		cfg.MaxWindowSize = 6
//	}
//
//	m := &MongoMemory{
//		client:        cfg.Client,
//		db:            cfg.DB,
//		collection:    cfg.DB.Collection(collectionName),
//		maxWindowSize: cfg.MaxWindowSize,
//		ctx:           cfg.Ctx,
//	}
//
//	go m.backgroundSync()
//	return m
//}
//
//func GetDefaultMemory(ctx context.Context) *MongoMemory {
//	return NewMongoMemory(MongoMemoryConfig{
//		Client:        dalmongo.Client,
//		DB:            dalmongo.DB,
//		MaxWindowSize: 6,
//		Ctx:           ctx,
//	})
//}
//
//func (m *MongoMemory) backgroundSync() {
//	ticker := time.NewTicker(5 * time.Minute)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			m.syncMemoryToDB()
//		case <-m.ctx.Done():
//			return
//		}
//	}
//}
//
//func (m *MongoMemory) syncMemoryToDB() {
//	batch := make([]mongo.WriteModel, 0, defaultBatchSize)
//
//	m.conversations.Range(func(key, value interface{}) bool {
//		conv := value.(*Conversation)
//		conv.mu.Lock()
//		defer conv.mu.Unlock()
//
//		update := mongo.NewUpdateOneModel().
//			SetFilter(bson.M{"_id": conv.ID}).
//			SetUpdate(bson.M{"$set": bson.M{"messages": conv.Messages}}).
//			SetUpsert(true)
//
//		batch = append(batch, update)
//
//		if len(batch) >= defaultBatchSize {
//			m.bulkWrite(batch)
//			batch = batch[:0]
//		}
//		return true
//	})
//
//	if len(batch) > 0 {
//		m.bulkWrite(batch)
//	}
//}
//
//func (m *MongoMemory) bulkWrite(models []mongo.WriteModel) {
//	opts := options.BulkWrite().SetOrdered(false)
//	_, err := m.collection.BulkWrite(context.Background(), models, opts)
//	if err != nil {
//		fmt.Printf("Bulk write error: %v\n", err)
//	}
//}
//
//func (m *MongoMemory) GetConversation(id string, createIfNotExist bool) *Conversation {
//	if val, ok := m.conversations.Load(id); ok {
//		return val.(*Conversation)
//	}
//
//	var convDoc struct {
//		ID       string            `bson:"_id"`
//		Messages []*schema.Message `bson:"messages"`
//	}
//
//	filter := bson.M{"_id": id}
//	err := m.collection.FindOne(m.ctx, filter).Decode(&convDoc)
//
//	switch {
//	case err == nil:
//		conv := &Conversation{
//			ID:            convDoc.ID,
//			Messages:      convDoc.Messages,
//			maxWindowSize: m.maxWindowSize,
//			collection:    m.collection,
//		}
//		m.conversations.Store(id, conv)
//		return conv
//
//	case err == mongo.ErrNoDocuments && createIfNotExist:
//		newConv := &Conversation{
//			ID:            id,
//			Messages:      make([]*schema.Message, 0),
//			maxWindowSize: m.maxWindowSize,
//			collection:    m.collection,
//		}
//		m.conversations.Store(id, newConv)
//		return newConv
//
//	default:
//		return nil
//	}
//}
//
//func (m *MongoMemory) ListConversations() ([]string, error) {
//	opts := options.Find().
//		SetProjection(bson.M{"_id": 1}).
//		SetBatchSize(1000)
//
//	cursor, err := m.collection.Find(m.ctx, bson.D{}, opts)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(m.ctx)
//
//	var ids []string
//	for cursor.Next(m.ctx) {
//		var result struct {
//			ID string `bson:"_id"`
//		}
//		if err := cursor.Decode(&result); err != nil {
//			continue
//		}
//		ids = append(ids, result.ID)
//	}
//	return ids, nil
//}
//
//func (m *MongoMemory) DeleteConversation(id string) error {
//	m.conversations.Delete(id)
//
//	_, err := m.collection.DeleteOne(m.ctx, bson.M{"_id": id})
//	return err
//}
//
//func (c *Conversation) Append(msg *schema.Message) error {
//	//c.mu.Lock()
//	//defer c.mu.Unlock()
//
//	c.Messages = append(c.Messages, msg)
//
//	go func() {
//		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//		defer cancel()
//
//		update := bson.M{
//			"$push": bson.M{
//				"messages": bson.M{
//					"$each":     []*schema.Message{msg},
//					"$position": 0,
//				},
//			},
//		}
//		_, err := c.collection.UpdateByID(ctx, c.ID, update)
//		if err != nil {
//			fmt.Printf("Failed to update conversation %s: %v\n", c.ID, err)
//		}
//	}()
//
//	return nil
//}
//
//func (c *Conversation) GetFullMessages() []*schema.Message {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	return c.Messages
//}
//
//func (c *Conversation) GetMessages() []*schema.Message {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	start := 0
//	if len(c.Messages) > c.maxWindowSize {
//		start = len(c.Messages) - c.maxWindowSize
//	}
//	return c.Messages[start:]
//}
