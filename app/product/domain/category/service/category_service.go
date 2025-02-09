package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	categoryrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/repository"
)

type CategoryService struct{}

var categoryServiceIns = &CategoryService{}

func GetCategoryService() *CategoryService {
	return categoryServiceIns
}

func (s *CategoryService) GetCategoryById(ctx context.Context, categoryId uint32) (*entity.CategoryEntity, error) {
	return nil, nil
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]*entity.CategoryEntity, error) {
	return nil, nil
}

func (s *CategoryService) BatchGetCategories(ctx context.Context, categoryIds []uint32) ([]*entity.CategoryEntity, error) {
	categories, err := categoryrepo.GetFactory().GetCategoryRepository().BatchGetCategories(ctx, categoryIds)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
