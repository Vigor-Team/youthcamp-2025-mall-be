package application

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type IncrStockService struct {
	ctx context.Context
} // NewIncrStockService new IncrStockService
func NewIncrStockService(ctx context.Context) *IncrStockService {
	return &IncrStockService{ctx: ctx}
}

// Run create note info
func (s *IncrStockService) Run(req *product.IncrStockReq) (resp *product.IncrStockResp, err error) {
	// Finish your business logic.

	return
}
