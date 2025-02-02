package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
)

type GetConversationIdService struct {
	ctx context.Context
} // NewGetConversationIdService new GetConversationIdService
func NewGetConversationIdService(ctx context.Context) *GetConversationIdService {
	return &GetConversationIdService{ctx: ctx}
}

// Run create note info
func (s *GetConversationIdService) Run(req *llm.GetConversationIdRequest) (resp *llm.GetConversationIdResponse, err error) {

	return
}
