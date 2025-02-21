package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/strategy"
)

type ProductRepository interface {
	AddProduct(ctx context.Context, product *entity.ProductEntity) error
	UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error
	GetProductById(ctx context.Context, productId uint32) (*entity.ProductEntity, error)
	ListProducts(ctx context.Context, strategy strategy.ListingStrategy) ([]*entity.ProductEntity, error)
	BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error)
	SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error)
}
