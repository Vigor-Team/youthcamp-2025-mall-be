package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestUpdateCart_Run(t *testing.T) {
	//// todo: edit your unit test
	dal.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewUpdateCartService(ctx)
	req := &cart.UpdateCartReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 3560968192,
			Quantity:  3,
		},
	}
	_, err := s.Run(req)
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
