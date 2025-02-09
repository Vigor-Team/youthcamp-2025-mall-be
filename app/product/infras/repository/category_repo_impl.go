package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/po"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (c *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, categoryId int64) (*entity.CategoryEntity, error) {
	return nil, nil
}

func (c *CategoryRepositoryImpl) GetCategories(ctx context.Context) ([]*entity.CategoryEntity, error) {
	return nil, nil
}

func (c *CategoryRepositoryImpl) BatchGetCategories(ctx context.Context, categoryIds []uint32) ([]*entity.CategoryEntity, error) {
	categories := []po.Category{}
	err := c.db.Where("id IN ?", categoryIds).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	ret := make([]*entity.CategoryEntity, 0, len(categories))
	for _, v := range categories {
		ret = append(ret, &entity.CategoryEntity{
			ID:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return ret, nil
}
