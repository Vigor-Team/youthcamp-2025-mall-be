package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
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

func TestAddSeckillProduct(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	// 加载商品
	key := redis.GetProductStockKey(2629832704)
	_, err := redis.RedisClient.Set(ctx, key, 100, 0).Result()
	if err != nil {
		t.Fatalf("set product stock failed: %v", err)
	}

}
