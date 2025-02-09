package application

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type OfflineProductService struct {
	ctx context.Context
} // NewOfflineProductService new OfflineProductService
func NewOfflineProductService(ctx context.Context) *OfflineProductService {
	return &OfflineProductService{ctx: ctx}
}

// Run create note info
func (s *OfflineProductService) Run(req *product.OfflineProductReq) (resp *product.OfflineProductResp, err error) {
	// Finish your business logic.

	return
}
