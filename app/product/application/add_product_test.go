package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
)

func TestAddProduct_Run(t *testing.T) {
	repository.Init()
	ctx := context.Background()
	s := NewAddProductService(ctx)
	// init req and assert value

	req := &product.AddProductReq{
		Name:        "test",
		Description: "test",
		Price:       1,
		Stock:       1,
		SpuName:     "test",
		SpuPrice:    1,
		Picture:     "test",
		CategoryIds: []uint32{1, 2},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
