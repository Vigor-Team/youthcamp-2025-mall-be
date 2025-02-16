package service

import (
	"context"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"testing"
)

func TestSeckillPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSeckillPlaceOrderService(ctx)
	// init req and assert value

	req := &order.SeckillPlaceOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
