package repository

import (
	"context"
	"errors"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository/converter"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (c *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, categoryId uint32) (*entity.CategoryEntity, error) {
	categories := make([]*po.Category, 0)
	if err := c.db.WithContext(ctx).Where("id = ?", categoryId).Find(&categories).Error; err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, errors.New("category not found")
	}
	do, err := converter.CategoryPOWithDOConverter.Convert2DO(ctx, categories[0])
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (c *CategoryRepositoryImpl) GetCategories(ctx context.Context) ([]*entity.CategoryEntity, error) {
	categories := make([]*po.Category, 0)
	if err := c.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	ret := make([]*entity.CategoryEntity, 0, len(categories))
	for _, v := range categories {
		do, err := converter.CategoryPOWithDOConverter.Convert2DO(ctx, v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, do)
	}
	return ret, nil
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
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}
	return ret, nil
}
