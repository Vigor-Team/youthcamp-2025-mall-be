package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	productrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/repository"
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

func (s *ProductQueryService) ListProducts(ctx context.Context, categoryId string) ([]*entity.ProductEntity, error) {
	filterParam := make(map[string]interface{}, 1)
	if categoryId != "" {
		filterParam["category_id"] = categoryId
	} else {
		filterParam = nil
	}
	products, err := productrepo.GetFactory().GetProductRepository().ListProducts(ctx, filterParam)
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
