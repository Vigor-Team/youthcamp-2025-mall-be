package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"testing"
)

func TestSeckillPlaceOrder_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewSeckillPlaceOrderService(ctx)
	// init req and assert value

	req := &order.SeckillPlaceOrderReq{
		UserId: 1,
		TempId: 2618155008,
		Email:  "1231231",
		Address: &order.Address{
			StreetAddress: "1231",
			City:          "1231",
			State:         "123",
			Country:       "123",
			ZipCode:       1,
		},
		UserCurrency: "12312",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestAddSeckillProduct(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	key := redis.GetProductStockKey(2629832704)
	_, err := redis.RedisClient.Set(ctx, key, 100, 0).Result()
	if err != nil {
		t.Fatalf("set product stock failed: %v", err)
	}

}
