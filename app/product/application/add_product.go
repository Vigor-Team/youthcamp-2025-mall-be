package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"

	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	//productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
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
	categories, err := categoryservice.GetCategoryService().BatchGetCategories(context.Background(), req.CategoryIds)
	if err != nil {
		return nil, err
	}
	entity, err := converter.ConvertAddReq2Entity(s.ctx, req)
	if err != nil {
		return nil, err
	}
	entity.Categories = categories
	err = productservice.GetProductUpdateService().AddProduct(context.Background(), entity)
	return
}
