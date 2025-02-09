package application

import (
	"context"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type DecrStockService struct {
	ctx context.Context
} // NewDecrStockService new DecrStockService
func NewDecrStockService(ctx context.Context) *DecrStockService {
	return &DecrStockService{ctx: ctx}
}

// Run create note info
func (s *DecrStockService) Run(req *product.DecrStockReq) (resp *product.DecrStockResp, err error) {
	err = productservice.GetProductStockService().DecreaseProductStock(context.Background(), req.GetId(), req.GetDecr())
	if err != nil {
		return nil, err
	}
	return
}
