package conversation

import (
	"context"
	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	"github.com/cloudwego/eino/schema"
	"testing"
)

func TestConvRun(t *testing.T) {
	dalmongo.Init()
	t.Log("TestMongoMemTest")
	bucket := GetDefaultBucket(context.Background())
	conv, err := bucket.GetConversation("111", "222")
	if err != nil {
		t.Log("GetConversation err")
	}
	if conv == nil {
		t.Error("GetConversation nil")
		return
	}
	_ = conv.Append(&schema.Message{
		Role:    "User",
		Content: "hello",
	})
	_ = conv.Append(&schema.Message{
		Role:    "Assistant",
		Content: "world",
	})
	messages := conv.GetFullMessages()
	t.Log(messages)
	getMessages := conv.GetMessages()
	t.Log(getMessages)
	err = bucket.DeleteConversation("111", "222")
	if err != nil {
		t.Error("DeleteConversation err")
		return
	}
}
