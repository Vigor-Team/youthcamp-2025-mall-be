package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
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
		return nil, err
	}

	categories, err := categoryservice.GetCategoryService().BatchGetCategories(s.ctx, req.GetCategoryIds())
	if err != nil {
		return nil, err
	}

	target, err := converter.ConvertUpdateReq2Entity(s.ctx, req)
	if err != nil {
		return nil, err
	}
	target.Categories = categories

	err = updateService.UpdateProduct(s.ctx, origin, target)
	if err != nil {
		return nil, err
	}
	return
}
