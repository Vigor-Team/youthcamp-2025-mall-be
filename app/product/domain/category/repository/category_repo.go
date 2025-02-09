package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, categoryId int64) (*entity.CategoryEntity, error)
	GetCategories(ctx context.Context) ([]*entity.CategoryEntity, error)
	BatchGetCategories(ctx context.Context, categoryIds []uint32) ([]*entity.CategoryEntity, error)
}
