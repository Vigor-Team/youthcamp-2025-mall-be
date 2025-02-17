package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"testing"
)

func TestSeckillPrePlaceOrder_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	mq.StartConsumer(ctx, mq.Client, 10)
	_ = redis.InitSnowflake()

	s := NewSeckillPrePlaceOrderService(ctx)
	// init req and assert value

	req := &order.SeckillPrePlaceOrderReq{
		UserId:    1,
		ProductId: 2629832704,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}

func TestSeckillPlaceOrderRun(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	preOrderKey := redis.GetOrderPreOrderKey(3851280384)
	productOrderKey := redis.GetProductOrderKey(2629832704)

	if err := redis.RedisClient.SRem(ctx, productOrderKey, 1).Err(); err != nil {
		t.Fatalf("remove product order failed: %v", err)
	}
	if err := redis.RedisClient.HDel(ctx, preOrderKey, "user_id", "product_id").Err(); err != nil {
		t.Fatalf("remove pre order failed: %v", err)
	}

}
