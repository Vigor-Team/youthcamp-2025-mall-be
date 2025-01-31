package chat

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

var (
	chatModel model.ChatModel
	once      sync.Once
)

func InitChatModel(ctx context.Context) {
	once.Do(func() {
		chatModel = createOpenAIChatModel(ctx)
	})
}

// todo 历史消息
func GenerateLlmResponse(ctx context.Context, in string) *schema.Message {
	messages := createMessagesFromTemplate(in)
	return generate(ctx, chatModel, messages)
}

func StreamLlmResponse(ctx context.Context, in string) *schema.StreamReader[*schema.Message] {
	messages := createMessagesFromTemplate(in)
	return stream(ctx, chatModel, messages)
}

func createMessagesFromTemplate(message string) []*schema.Message {
	template := createTemplate()

	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "电商客服",
		"style":    "友好、耐心、专业、关注用户体验",
		"object":   "帮助用户解决问题包括但不限于商品搜索，帮助下单，售后处理等等，提升用户体验。",
		"question": message,
	})
	if err != nil {
		klog.CtxErrorf(context.Background(), "format template failed: %v\n", err)
	}
	return messages
}

func createTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是{object}"),

		//schema.MessagesPlaceholder("chat_history", true),

		schema.UserMessage("问题: {question}"),
	)
}

func createOpenAIChatModel(ctx context.Context) model.ChatModel {
	//key := os.Getenv("OPENAI_API_KEY")
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   "gpt-3.5-turbo",
		APIKey:  "sk-H6HQY3ZQ5e2xIcGiHEHtoa4MudvS3TzO6DQfrrgCxNqgy7T4",
		BaseURL: "https://api.chatanywhere.tech",
	})
	if err != nil {
		klog.CtxFatalf(ctx, "create openai chat model failed, err=%v", err)
	}
	return chatModel
}

func generate(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.Message {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		klog.CtxErrorf(ctx, "chat generate failed: %v", err)
	}
	return result
}

func stream(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
	if err != nil {
		klog.CtxErrorf(ctx, "chat stream failed: %v", err)
	}
	return result
}
