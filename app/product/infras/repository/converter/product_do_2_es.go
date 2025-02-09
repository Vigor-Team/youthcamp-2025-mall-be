package converter

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
)

type productDo2ESConverter struct{}

var ProductDo2ESConverter = &productDo2ESConverter{}

func (c *productDo2ESConverter) Convert2Es(_ context.Context, product *entity.ProductEntity) *entity.ProductES {
	categoryNames := make([]string, len(product.Categories))
	for i, category := range product.Categories {
		categoryNames[i] = category.Name
	}
	return &entity.ProductES{
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
