package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
)

type ProductQueryService struct {
}

var productQueryServiceIns = &ProductQueryService{}

func GetProductQueryService() *ProductQueryService {
	return productQueryServiceIns
}

func (s *ProductQueryService) GetProductById(ctx context.Context, productId uint32) (*entity.ProductEntity, error) {
	return nil, nil
}

func (s *ProductQueryService) ListProducts(ctx context.Context, filterParam map[string]interface{}) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (s *ProductQueryService) SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (s *ProductQueryService) BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	return nil, nil
}
