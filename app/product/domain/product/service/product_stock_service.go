package service

import (
	"context"
	productrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/repository"
)

type ProductStockService struct {
}

var productStockServiceIns = &ProductStockService{}

func GetProductStockService() *ProductStockService {
	return productStockServiceIns
}

func (s *ProductStockService) IncreaseProductStock(ctx context.Context, productId, incr uint32) error {
	return productrepo.GetFactory().GetStockRepository().IncrStock(ctx, productId, incr)
}

func (s *ProductStockService) DecreaseProductStock(ctx context.Context, productId, decr uint32) error {
	return productrepo.GetFactory().GetStockRepository().DecrStock(ctx, productId, decr)
}
