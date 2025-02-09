package converter

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	po2 "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
)

type productDO2POConverter struct{}

var ProductDO2POConverter = &productDO2POConverter{}

func (c *productDO2POConverter) Convert2po(_ context.Context, product *entity.ProductEntity) (*po2.Product, error) {
	categories := make([]po2.Category, 0, len(product.Categories))
	for _, category := range product.Categories {
		categories = append(categories, po2.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	ret := &po2.Product{
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
	return ret, nil
}
