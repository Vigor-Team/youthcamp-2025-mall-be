package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
)

type ProductUpdateService struct{}

var productUpdateService ProductUpdateService

func GetProductUpdateServiceInstance() *ProductUpdateService {
	return &productUpdateService
}

func (s *ProductUpdateService) AddProduct(ctx context.Context, product *entity.ProductEntity) error {
	return nil
}

func (s *ProductUpdateService) UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	return nil
}
