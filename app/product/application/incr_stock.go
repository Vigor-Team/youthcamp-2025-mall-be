package application

import (
	"context"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
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
	err = productservice.GetProductStockService().IncreaseProductStock(context.Background(), req.GetId(), req.GetIncr())
	if err != nil {
		return nil, err
	}
	return
}
