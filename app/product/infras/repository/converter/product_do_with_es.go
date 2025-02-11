package converter

import (
	"context"
	entity2 "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
)

type productDoWithESConverter struct{}

var ProductDoWithESConverter = &productDoWithESConverter{}

func (c *productDoWithESConverter) Convert2ES(_ context.Context, product *entity2.ProductEntity) *entity2.ProductES {
	categoryNames := make([]string, len(product.Categories))
	for i, category := range product.Categories {
		categoryNames[i] = category.Name
	}
	return &entity2.ProductES{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		Picture:       product.Picture,
		SpuName:       product.SpuName,
		SpuPrice:      product.SpuPrice,
		Price:         product.Price,
		Stock:         product.Stock,
		Status:        product.Status,
		CategoryNames: categoryNames,
	}
}

func (c *productDoWithESConverter) Convert2DO(_ context.Context, product *entity2.ProductES) *entity2.ProductEntity {
	categories := make([]*entity2.CategoryEntity, 0, len(product.CategoryNames))
	for _, categoryName := range product.CategoryNames {
		categories = append(categories, &entity2.CategoryEntity{
			Name: categoryName,
		})
	}
	return &entity2.ProductEntity{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Status:      product.Status,
		SpuName:     product.SpuName,
		SpuPrice:    product.SpuPrice,
		Picture:     product.Picture,
		Categories:  categories,
	}
}
