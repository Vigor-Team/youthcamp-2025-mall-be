package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	get, err := productservice.GetProductQueryService().SearchProducts(s.ctx, req.Query)
	if err != nil {
		return nil, err
	}
	products := make([]*product.Product, 0, len(get))
	for _, v := range get {
		products = append(products, converter.ProductConvertEntity2DTO(v))
	}
	resp = &product.SearchProductsResp{
		Results: products,
	}
	return
}
