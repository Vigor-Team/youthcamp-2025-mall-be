package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
)

func TestUpdateProduct_Run(t *testing.T) {
	repository.Init()
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &product.UpdateProductReq{
		Id:          2629832704,
		Name:        "Sports Shoes",
		Description: "Comfortable sports shoes, suitable for running and everyday wear, breathable design.",
		Price:       400,
		Stock:       100,
		SpuName:     "Footwear Series",
		SpuPrice:    399,
		Picture:     "https://example.com/sportsshoes.jpg",
		CategoryIds: []uint32{2, 3},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
