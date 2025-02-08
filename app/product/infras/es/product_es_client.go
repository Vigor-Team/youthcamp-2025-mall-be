package es

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/entity"
)

type ProductESClient struct{}

var productESClient ProductESClient

func GetProductESClient() *ProductESClient {
	return &productESClient
}

func (c *ProductESClient) BatchGetProductById(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (c *ProductESClient) SearchProduct(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func getEntityFormSource(source string) *entity.ProductEntity {
	return nil
}

func getDocFromEntity(entity *entity.ProductEntity) map[string]interface{} {
	return nil
}
