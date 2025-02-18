package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"testing"
)

func TestQueryOrder_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewQueryOrderService(ctx)
	// init req and assert value

	req := &order.QueryOrderReq{
		UserId:  1,
		OrderId: "3108433920",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
