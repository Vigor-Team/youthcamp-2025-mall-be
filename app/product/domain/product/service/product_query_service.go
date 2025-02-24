package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	productrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/repository"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/strategy"
)

type ProductQueryService struct {
}

var productQueryServiceIns = &ProductQueryService{}

func GetProductQueryService() *ProductQueryService {
	return productQueryServiceIns
}

func (s *ProductQueryService) GetProductById(ctx context.Context, productId uint32) (*entity.ProductEntity, error) {
	id, err := productrepo.GetFactory().GetProductRepository().GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *ProductQueryService) ListProducts(ctx context.Context, categoryId uint32, role string) ([]*entity.ProductEntity, error) {
	filterParam := make(map[string]interface{}, 1)
	if categoryId != 0 {
		filterParam["category_id"] = categoryId
	} else {
		filterParam = nil
	}
	listingStrategy := strategy.NewListingStrategy(role, filterParam)
	products, err := productrepo.GetFactory().GetProductRepository().ListProducts(ctx, listingStrategy)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductQueryService) SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	products, err := productrepo.GetFactory().GetProductRepository().SearchProducts(ctx, keyword)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductQueryService) BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	products, err := productrepo.GetFactory().GetProductRepository().BatchGetProducts(ctx, productIds)
	if err != nil {
		return nil, err
	}
	return products, nil
}
