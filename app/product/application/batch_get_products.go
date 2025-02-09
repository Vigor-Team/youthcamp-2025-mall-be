package application

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type BatchGetProductsService struct {
	ctx context.Context
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	// Finish your business logic.

	return
}
