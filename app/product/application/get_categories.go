package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	categoryservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type GetCategoriesService struct {
	ctx context.Context
} // NewGetCategoriesService new GetCategoriesService
func NewGetCategoriesService(ctx context.Context) *GetCategoriesService {
	return &GetCategoriesService{ctx: ctx}
}

// Run create note info
func (s *GetCategoriesService) Run(req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	get, err := categoryservice.GetCategoryService().GetCategories(s.ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*product.Category, 0, len(get))
	for _, v := range get {
		categories = append(categories, converter.CategoryConvertEntity2DTO(v))
	}
	resp = &product.GetCategoriesResp{
		Categories: categories,
	}
	return
}
