package converter

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
)

type categoryPOWithDOConverter struct{}

var CategoryPOWithDOConverter = &categoryPOWithDOConverter{}

func (c *categoryPOWithDOConverter) Convert2PO(ctx context.Context, category *entity.CategoryEntity) (*po.Category, error) {
	ret := &po.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return ret, nil
}

func (c *categoryPOWithDOConverter) Convert2DO(ctx context.Context, category *po.Category) (*entity.CategoryEntity, error) {
	ret := &entity.CategoryEntity{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return ret, nil
}
