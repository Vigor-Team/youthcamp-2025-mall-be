package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/entity"
)

type ProductRepositoryImpl struct{}

func (p *ProductRepositoryImpl) AddProduct(ctx context.Context, product *entity.ProductEntity) error {
	return nil
}

func (p *ProductRepositoryImpl) UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	return nil
}

func (p *ProductRepositoryImpl) GetProductById(ctx context.Context, productId int64) (*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) ListProducts(ctx context.Context, filterParam map[string]interface{}) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	return nil, nil
}
