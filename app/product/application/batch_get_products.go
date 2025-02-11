package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/converter"
	productservice "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type BatchGetProductsService struct {
	ctx context.Context
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	batch, err := productservice.GetProductQueryService().BatchGetProducts(context.Background(), req.GetIds())
	products := map[uint32]*product.Product{}
	for _, p := range batch {
		products[p.ID] = converter.ProductConvertEntity2DTO(p)
	}
	resp = &product.BatchGetProductsResp{
		Products: products,
	}
	return
}
