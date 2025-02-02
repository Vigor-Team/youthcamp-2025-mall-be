package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
)

type GetHistoryService struct {
	ctx context.Context
} // NewGetHistoryService new GetHistoryService
func NewGetHistoryService(ctx context.Context) *GetHistoryService {
	return &GetHistoryService{ctx: ctx}
}

// Run create note info
func (s *GetHistoryService) Run(req *llm.GetHistoryRequest) (resp *llm.GetHistoryResponse, err error) {
	// Finish your business logic.

	return
}
