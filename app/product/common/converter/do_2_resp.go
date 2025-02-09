package converter

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

func ProductConvertEntity2DTO(e *entity.ProductEntity) *product.Product {
	categoryNames := make([]string, 0, len(e.Categories))
	for _, c := range e.Categories {
		categoryNames = append(categoryNames, c.Name)
	}
	return &product.Product{
		Id:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
		SpuName:     e.SpuName,
		SpuPrice:    e.SpuPrice,
		Picture:     e.Picture,
		Status:      product.Status(e.Status),
		Categories:  categoryNames,
	}
}

func CategoryConvertEntity2DTO(e *entity.CategoryEntity) *product.Category {
	return &product.Category{
		Id:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}
