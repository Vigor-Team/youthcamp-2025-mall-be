package converter

import (
	"context"
	entity2 "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
)

type productPO2DOConverter struct{}

var ProductPO2DOConverter = productPO2DOConverter{}

func (converter *productPO2DOConverter) Convert2do(_ context.Context, po *po.Product) (*entity2.ProductEntity, error) {
	categories := make([]*entity2.CategoryEntity, 0, len(po.Categories))
	for _, category := range po.Categories {
		categories = append(categories, &entity2.CategoryEntity{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	do := &entity2.ProductEntity{
		ID:          po.ID,
		Name:        po.Name,
		Price:       po.Price,
		Description: po.Description,
		Stock:       po.Stock,
		Status:      po.Status,
		SpuName:     po.SpuName,
		SpuPrice:    po.SpuPrice,
		Picture:     po.Picture,
		Categories:  categories,
	}
	return do, nil
}
