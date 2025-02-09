package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
)

type CategoryRepositoryImpl struct {
}

func (c *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, categoryId int64) (*entity.CategoryEntity, error) {
	return nil, nil
}

func (c *CategoryRepositoryImpl) GetCategories(ctx context.Context) ([]*entity.CategoryEntity, error) {
	return nil, nil
}

func (c *CategoryRepositoryImpl) BatchGetCategories(ctx context.Context, categoryIds []uint32) ([]*entity.CategoryEntity, error) {
	return nil, nil
}
