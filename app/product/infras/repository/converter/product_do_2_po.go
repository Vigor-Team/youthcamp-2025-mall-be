package converter

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/po"
)

type productDO2POConverter struct{}

var ProductDO2POConverter = &productDO2POConverter{}

func (c *productDO2POConverter) Convert2po(_ context.Context, product *entity.ProductEntity) (*po.Product, error) {
	categories := make([]po.Category, 0, len(product.Categories))
	for _, category := range product.Categories {
		categories = append(categories, po.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	ret := &po.Product{
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
