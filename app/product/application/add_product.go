package application

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	categories, err := categoryservice.GetCategoryService().BatchGetCategories(s.ctx, req.CategoryIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "BatchGetCategories failed, err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetCategory, "BatchGetCategories failed")
	}
	entity, err := converter.ConvertAddReq2Entity(s.ctx, req)
	if err != nil {
		klog.CtxErrorf(s.ctx, "ConvertAddReq2Entity failed, err:%v", err)
		return nil, kerrors.NewBizStatusError(errno.ErrInternal, "ConvertAddReq2Entity failed")
	}
	entity.Categories = categories
	err = productservice.GetProductUpdateService().AddProduct(s.ctx, entity)
	if err != nil {
		klog.CtxErrorf(s.ctx, "AddProduct failed, err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrAddProduct, "AddProduct failed")
	}
	return
}
