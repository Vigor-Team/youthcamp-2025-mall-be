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
	reqs := []*product.AddProductReq{}
	reqs = append(reqs, &product.AddProductReq{
		Name:        "Wireless Bluetooth Earphones",
		Description: "High-quality wireless Bluetooth earphones, clear sound, long battery life.",
		Price:       299,
		Stock:       50,
		SpuName:     "Bluetooth Earphone Series",
		SpuPrice:    299,
		Picture:     "https://example.com/earphone.jpg",
		CategoryIds: []uint32{1, 2},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "Smartwatch",
		Description: "Multi-functional smartwatch with heart rate monitoring, GPS, and waterproof design.",
		Price:       899,
		Stock:       30,
		SpuName:     "Watch Series",
		SpuPrice:    899,
		Picture:     "https://example.com/smartwatch.jpg",
		CategoryIds: []uint32{1, 3},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "Sports Shoes",
		Description: "Comfortable sports shoes, suitable for running and everyday wear, breathable design.",
		Price:       399,
		Stock:       100,
		SpuName:     "Footwear Series",
		SpuPrice:    399,
		Picture:     "https://example.com/sportsshoes.jpg",
		CategoryIds: []uint32{2, 3},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "Leather Wallet",
		Description: "Premium leather wallet with multiple compartments, stylish and durable.",
		Price:       129,
		Stock:       200,
		SpuName:     "Wallet Series",
		SpuPrice:    129,
		Picture:     "https://example.com/wallet.jpg",
		CategoryIds: []uint32{1, 3},
	})

	for _, req := range reqs {
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}
}
