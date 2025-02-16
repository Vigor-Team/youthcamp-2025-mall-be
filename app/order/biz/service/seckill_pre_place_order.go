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

	script, err := script.GetPreSeckillScript()
	if err != nil {
		return nil, err
	}

	// todo 生成tempId 分布式id
	preOrderId, err := redis.NextId(s.ctx, "pre_order_id")

	productStockKey := redis.GetProductStockKey(productId)
	tempKey := redis.GetSeckillTempKey(strconv.Itoa(int(preOrderId)))

	// todo 临时过期时间
	expireSeconds := 10 * 60

	result, err := redis.RedisClient.Eval(s.ctx, script, []string{productStockKey, tempKey}, userId, productId, expireSeconds).Result()
	if err != nil {
		return nil, err
	}

	if resultMap, ok := result.(map[string]interface{}); ok {
		if errMsg, exists := resultMap["err"]; exists {
			switch errMsg {
			case "OUT_OF_STOCK":
				return nil, kerrors.NewBizStatusError(1, "out of stock")
			case "DUPLICATE_USER":
				return nil, kerrors.NewBizStatusError(2, "please do not repeat the operation")
			}
		}
		// todo 发送预占消息到 RabbitMQ，失败则回滚
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
			TempId: strconv.Itoa(int(preOrderId)),
		}
	}
	return nil, kerrors.NewBizStatusError(3, "system error")
}
