package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DecrStockService struct {
	ctx context.Context
} // NewDecrStockService new DecrStockService
func NewDecrStockService(ctx context.Context) *DecrStockService {
	return &DecrStockService{ctx: ctx}
}

// Run create note info
func (s *DecrStockService) Run(req *product.DecrStockReq) (resp *product.DecrStockResp, err error) {
	err = productservice.GetProductStockService().DecreaseProductStock(s.ctx, req.GetId(), req.GetDecr())
	if err != nil {
		klog.CtxErrorf(s.ctx, "productservice.GetProductStockService().DecreaseProductStock.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrDecrProductStock, "DecreaseProductStock failed")
	}
	return
}
