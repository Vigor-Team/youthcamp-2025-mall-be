package chat

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"log"
)

// todo 历史消息
func GenerateLlmResponse(ctx context.Context, in string) *schema.Message {
	messages := createMessagesFromTemplate(in)
	llm := createOpenAIChatModel(ctx)
	return generate(ctx, llm, messages)
}

func StreamLlmResponse(ctx context.Context, in string) *schema.StreamReader[*schema.Message] {
	messages := createMessagesFromTemplate(in)
	llm := createOpenAIChatModel(ctx)
	return stream(ctx, llm, messages)
}

func createMessagesFromTemplate(message string) []*schema.Message {
	template := createTemplate()

	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "程序员鼓励师",
		"style":    "积极、温暖且专业",
		"question": message,
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
			schema.UserMessage("我觉得自己写的代码太烂了"),
			schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
		},
	})
	if err != nil {
		log.Fatalf("format template failed: %v\n", err)
	}
	return messages
}

func createTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)
}

func createOpenAIChatModel(ctx context.Context) model.ChatModel {
	//key := os.Getenv("OPENAI_API_KEY")
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   "gpt-4o",
		APIKey:  "sk-H6HQY3ZQ5e2xIcGiHEHtoa4MudvS3TzO6DQfrrgCxNqgy7T4",
		BaseURL: "https://api.chatanywhere.tech",
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}

func generate(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.Message {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		log.Fatalf("chat generate failed: %v", err)
	}
	return result
}

func stream(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
	if err != nil {
		log.Fatalf("chat generate failed: %v", err)
	}
	return result
}
