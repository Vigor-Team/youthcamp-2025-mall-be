package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestGetCart_Run(t *testing.T) {
	// dsn := "root:root@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local"
	dal.Init()
	// registryAddr = "127.0.0.1:8500"
	// serviceName = "cart"
	rpc.InitClient()
	ctx := context.Background()
	s := NewGetCartService(ctx)
	req := &cart.GetCartReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("resp: %v", resp)
}
