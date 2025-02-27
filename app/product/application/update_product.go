package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	updateService := productservice.GetProductUpdateService()
	queryService := productservice.GetProductQueryService()

	origin, err := queryService.GetProductById(s.ctx, req.GetId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "queryService.GetProductById.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetProduct, "GetProductById failed")
	}

	categories, err := categoryservice.GetCategoryService().BatchGetCategories(s.ctx, req.GetCategoryIds())
	if err != nil {
		klog.CtxErrorf(s.ctx, "BatchGetCategories.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetCategory, "BatchGetCategories failed")
	}

	target, err := converter.ConvertUpdateReq2Entity(s.ctx, req)
	if err != nil {
		klog.CtxErrorf(s.ctx, "ConvertUpdateReq2Entity.err:%v", err)
		return nil, kerrors.NewBizStatusError(errno.ErrInternal, "ConvertUpdateReq2Entity failed")
	}
	target.Categories = categories

	err = updateService.UpdateProduct(s.ctx, origin, target)
	if err != nil {
		klog.CtxErrorf(s.ctx, "updateService.UpdateProduct.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUpdateProduct, "UpdateProduct failed")
	}
	return
}
