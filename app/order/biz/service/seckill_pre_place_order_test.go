package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"
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
	var preOrder model.PreOrder
	if err := mysql.DB.Where("id = ?", 2668486656).First(&preOrder).Error; err != nil {
		t.Fatalf("get pre order failed: %v", err)
	}
	fmt.Println("preOrder: ", preOrder)

}
