package conversation

import (
	"context"
	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	"github.com/cloudwego/eino/schema"
	"testing"
)

func TestMongoMemTest(t *testing.T) {
	dalmongo.Init()
	t.Log("TestMongoMemTest")
	memory := GetDefaultMemory(context.Background())
	conversation := memory.GetConversation("111", true)
	if conversation == nil {
		t.Log("GetConversation nil")
	}
	conversation.Append(&schema.Message{
		Role:    "User",
		Content: "hello",
	})
	conversation.Append(&schema.Message{
		Role:    "Assistant",
		Content: "world",
	})
	messages := conversation.GetFullMessages()
	t.Log(messages)
	getMessages := conversation.GetMessages()
	t.Log(getMessages)
}
