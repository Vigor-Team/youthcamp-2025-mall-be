package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	get, err := productservice.GetProductQueryService().GetProductById(s.ctx, req.GetId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "productservice.GetProductQueryService().GetProductById.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetProduct, "GetProduct failed")
	}
	resp = &product.GetProductResp{
		Product: converter.ProductConvertEntity2DTO(get),
	}
	return
}
