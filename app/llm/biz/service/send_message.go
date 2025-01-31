package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
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
	resp = &llm.ChatResponse{
		//Response: chatresp.Content,
		Response: "hello",
	}
	return
}
