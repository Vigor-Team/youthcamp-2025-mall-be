package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCategoryService struct {
	ctx context.Context
} // NewGetCategoryService new GetCategoryService
func NewGetCategoryService(ctx context.Context) *GetCategoryService {
	return &GetCategoryService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryService) Run(req *product.GetCategoryReq) (resp *product.GetCategoryResp, err error) {
	get, err := categoryservice.GetCategoryService().GetCategoryById(s.ctx, req.GetId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "categoryservice.GetCategoryService().GetCategoryById.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetCategory, "GetCategory failed")
	}
	resp = &product.GetCategoryResp{
		Category: converter.CategoryConvertEntity2DTO(get),
	}
	return
}
