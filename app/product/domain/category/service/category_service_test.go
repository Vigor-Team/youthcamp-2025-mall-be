package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestBatchGetCategories(t *testing.T) {
	repository.Init()
	categories, err := GetCategoryService().BatchGetCategories(context.Background(), []uint32{1, 2, 3})
	assert.NotNil(t, categories)
	assert.Nil(t, err)
	for _, category := range categories {
		t.Logf("category: %v", category)
	}
}
