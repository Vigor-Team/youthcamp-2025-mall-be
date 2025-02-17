package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis/script"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"strconv"
	"time"
)

type SeckillPrePlaceOrderService struct {
	ctx context.Context
} // NewSeckillPrePlaceOrderService new SeckillPrePlaceOrderService

func NewSeckillPrePlaceOrderService(ctx context.Context) *SeckillPrePlaceOrderService {
	return &SeckillPrePlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *SeckillPrePlaceOrderService) Run(req *order.SeckillPrePlaceOrderReq) (resp *order.SeckillPrePlaceOrderResp, err error) {
	userId := req.UserId
	productId := req.ProductId
	if userId == 0 || productId == 0 {
		return nil, fmt.Errorf("invalid request")
	}

	// todo 限流检查

	seckillScript, err := script.GetPreSeckillScript()
	if err != nil {
		return nil, err
	}

	preOrderId, err := redis.NextId(s.ctx, redis.PreOrderNode)
	fmt.Println("preOrderId: ", preOrderId)

	productStockKey := redis.GetProductStockKey(productId)
	productOrderKey := redis.GetProductOrderKey(productId)
	preOrderKey := redis.GetOrderPreOrderKey(preOrderId)

	result, err := redis.RedisClient.Eval(s.ctx, seckillScript, []string{productStockKey, productOrderKey, preOrderKey}, userId, productId).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("result: ", result)

	if resultMap, ok := result.(map[string]interface{}); ok {
		if errMsg, exists := resultMap["err"]; exists {
			switch errMsg {
			case "OUT_OF_STOCK":
				return nil, kerrors.NewBizStatusError(1, "out of stock")
			case "DUPLICATE_USER":
				return nil, kerrors.NewBizStatusError(2, "please do not repeat the operation")
			}
		}
	}

	producer := mq.NewProducer(mq.Client)
	msg := mq.PreOrderMessage{
		TempID:    strconv.Itoa(int(preOrderId)),
		UserID:    userId,
		ProductID: productId,
		Timestamp: time.Now().Unix(),
	}
	if err := producer.PublishPreOrder(s.ctx, msg); err != nil {
		// todo 回滚
		return nil, err
	}

	resp = &order.SeckillPrePlaceOrderResp{
		TempId: preOrderId,
	}
	return
}
