package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.CtxErrorf(s.ctx, "productservice.GetProductStockService().IncreaseProductStock.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrIncrProductStock, "IncreaseProductStock failed")
	}
	return
}
