package application

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/embedding"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/utils"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
	"time"
)

func TestAddProduct_Run(t *testing.T) {
	repository.Init()
	ctx := context.Background()
	s := NewAddProductService(ctx)
	// init req and assert value
	reqs := []*product.AddProductReq{}
	reqs = append(reqs, &product.AddProductReq{
		Name:        "无线蓝牙耳机",
		Description: "高质量无线蓝牙耳机，音质清晰，电池续航长。",
		Price:       299,
		Stock:       50,
		SpuName:     "蓝牙耳机系列",
		SpuPrice:    299,
		Picture:     "https://example.com/earphone.jpg",
		CategoryIds: []uint32{1, 2},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "智能手表",
		Description: "多功能智能手表，支持心率监测，GPS定位，防水设计。",
		Price:       899,
		Stock:       30,
		SpuName:     "手表系列",
		SpuPrice:    899,
		Picture:     "https://example.com/smartwatch.jpg",
		CategoryIds: []uint32{1, 3},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "运动鞋",
		Description: "舒适的运动鞋，适合跑步和日常穿着，透气设计。",
		Price:       399,
		Stock:       100,
		SpuName:     "鞋类系列",
		SpuPrice:    399,
		Picture:     "https://example.com/sportsshoes.jpg",
		CategoryIds: []uint32{2, 3},
	})
	reqs = append(reqs, &product.AddProductReq{
		Name:        "皮革钱包",
		Description: "优质皮革钱包，多个隔层，时尚耐用。",
		Price:       129,
		Stock:       200,
		SpuName:     "钱包系列",
		SpuPrice:    129,
		Picture:     "https://example.com/wallet.jpg",
		CategoryIds: []uint32{1, 3},
	})

	for _, req := range reqs {
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}
	time.Sleep(5 * time.Second)
}

func TestEmbedding(t *testing.T) {
	eb, err := embedding.GetArkEmbedding(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetArkEmbedding err: %v", err)
	}
	texts := []string{
		"无线蓝牙耳机",
		"高质量无线蓝牙耳机，音质清晰，电池续航长。",
		"蓝牙耳机系列",
	}
	res, err := eb.EmbedStrings(context.Background(), texts)
	if err != nil {
		t.Fatalf("EmbedStrings err: %v", err)
	}
	var merged []float32
	if len(res) > 0 {
		merged = utils.MergeVectors(res)
	}
	fmt.Println("merged: ", merged)
}
