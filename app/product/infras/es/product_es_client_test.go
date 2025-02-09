package es

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestGetProductESClient(t *testing.T) {
	client := GetESClient()
	assert.NotNil(t, client)
}

func TestUpsertProductES(t *testing.T) {
	err := GetProductESClient().UpdateProduct(context.Background(), 1001, &entity.ProductES{
		ID:            1001,
		Name:          "Wireless Mechanical Keyboard",
		Description:   "A high-quality wireless mechanical keyboard with RGB lighting and hot-swappable switches.",
		Picture:       "https://example.com/images/keyboard.jpg",
		Price:         129.99,
		Stock:         50,
		SpuName:       "Wireless Keyboard Pro",
		SpuPrice:      149.99,
		Status:        1,
		CategoryNames: []string{"Electronics", "Keyboards"},
	})
	assert.Nil(t, err)

	err = GetProductESClient().UpdateProduct(context.Background(), 1002, &entity.ProductES{
		ID:            1002,
		Name:          "Gaming Mouse",
		Description:   "An ergonomic gaming mouse with 16000 DPI and customizable side buttons.",
		Picture:       "https://example.com/images/mouse.jpg",
		Price:         59.99,
		Stock:         30,
		SpuName:       "Gaming Mouse X",
		SpuPrice:      69.99,
		Status:        1,
		CategoryNames: []string{"Electronics", "Gaming Accessories"},
	})
	assert.Nil(t, err)

	err = GetProductESClient().UpdateProduct(context.Background(), 1003, &entity.ProductES{
		ID:            1003,
		Name:          "Noise-Canceling Headphones",
		Description:   "Over-ear noise-canceling headphones with high-fidelity sound and 30-hour battery life.",
		Picture:       "https://example.com/images/headphones.jpg",
		Price:         199.99,
		Stock:         20,
		SpuName:       "AudioPro NC700",
		SpuPrice:      219.99,
		Status:        1,
		CategoryNames: []string{"Electronics", "Audio"},
	})
	assert.Nil(t, err)
}

func TestBatchGetProductById(t *testing.T) {
	entities, err := GetProductESClient().BatchGetProductById(context.Background(), []uint32{1001, 1002, 1003})
	assert.Nil(t, err)
	for _, entity := range entities {
		fmt.Println(entity)
	}
}

func TestSearchProduct(t *testing.T) {
	entities, err := GetProductESClient().SearchProduct(context.Background(), "Over-ear")
	assert.Nil(t, err)
	for _, entity := range entities {
		fmt.Println(entity)
	}
}
