package service

import (
	"context"
	"testing"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
)

func TestDeleteItem_Run(t *testing.T) {
	dal.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewDeleteItemService(ctx)
	req := &cart.DeleteItemReq{
		UserId:    1,
		ProductId: 3560968192,
	}
	_, err := s.Run(req)
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
