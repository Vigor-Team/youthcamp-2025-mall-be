package main

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/service"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
)

// LlmServiceImpl implements the last service interface defined in the IDL.
type LlmServiceImpl struct{}

// SendMessage implements the LlmServiceImpl interface.
func (s *LlmServiceImpl) SendMessage(ctx context.Context, req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	resp, err = service.NewSendMessageService(ctx).Run(req)

	return resp, err
}

func (s *LlmServiceImpl) StreamMessage(req *llm.ChatRequest, stream llm.LlmService_StreamMessageServer) (err error) {
	ctx := context.Background()
	err = service.NewStreamMessageService(ctx).Run(req, stream)
	return
}
