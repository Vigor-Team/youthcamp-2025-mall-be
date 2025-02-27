package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type BatchGetProductsService struct {
	ctx context.Context
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	batch, err := productservice.GetProductQueryService().BatchGetProducts(s.ctx, req.GetIds())
	if err != nil {
		klog.CtxErrorf(s.ctx, "productservice.GetProductQueryService().BatchGetProducts.err:%v", err)
		err = kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "BatchGetProducts failed")
		return nil, err
	}
	products := map[uint32]*product.Product{}
	for _, p := range batch {
		products[p.ID] = converter.ProductConvertEntity2DTO(p)
	}
	resp = &product.BatchGetProductsResp{
		Products: products,
	}
	return
}
