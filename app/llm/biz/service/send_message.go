package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/mem"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type SendMessageService struct {
	ctx context.Context
} // NewSendMessageService new SendMessageService
func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

// Run create note info
func (s *SendMessageService) Run(req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	// Finish your business logic.
	//in := req.Message
	//chatresp := chat.GenerateLlmResponse(s.ctx, in)
	//resp = &llm.ChatResponse{
	//	Response: chatresp.Content,
	//}
	runnable, err := mallagent.BuildMallAgent(s.ctx, &mallagent.BuildConfig{MallAgent: &mallagent.MallAgentBuildConfig{}})
	if err != nil {
		klog.CtxErrorf(s.ctx, "build mall agent error: %v", err)
		return
	}
	userMessage := &mallagent.UserMessage{
		ChatId: uuid.New().String(),
		Query:  req.Message,
		UserId: req.UserId,
	}

	memory := mem.GetDefaultMemory()
	conversation := memory.GetConversation(userMessage.UserId, true)
	fmt.Println("conversation: ", conversation)
	rs, err := runnable.Invoke(s.ctx, userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "invoke mall agent error: %v", err)
		return
	}
	resp = &llm.ChatResponse{
		Response: rs.Content,
	}
	return
}
