package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
)

type DeleteHistoryService struct {
	ctx context.Context
} // NewDeleteHistoryService new DeleteHistoryService
func NewDeleteHistoryService(ctx context.Context) *DeleteHistoryService {
	return &DeleteHistoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteHistoryService) Run(req *llm.DeleteHistoryRequest) (resp *llm.DeleteHistoryResponse, err error) {
	// Finish your business logic.

	return
}
