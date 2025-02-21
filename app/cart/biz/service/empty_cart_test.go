package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestEmptyCart_Run(t *testing.T) {
	dal.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	req := &cart.EmptyCartReq{
		UserId: 1,
	}
	_, err := s.Run(req)
	if err != nil {
		t.Errorf("err: %v", err)
	}

}
