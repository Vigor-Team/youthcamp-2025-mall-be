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

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	get, err := productservice.GetProductQueryService().ListProducts(s.ctx, req.CategoryId, req.Role)
	if err != nil {
		klog.CtxErrorf(s.ctx, "productservice.GetProductQueryService().ListProducts.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetProduct, "ListProducts failed")
	}
	products := make([]*product.Product, 0, len(get))
	for _, v := range get {
		products = append(products, converter.ProductConvertEntity2DTO(v))
	}
	resp = &product.ListProductsResp{
		Products: products,
	}
	return
}
