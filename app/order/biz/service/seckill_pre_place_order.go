package service

import (
	"context"
	"strconv"
	"time"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis/script"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SeckillPrePlaceOrderService struct {
	ctx context.Context
}

// NewSeckillPrePlaceOrderService new SeckillPrePlaceOrderService
func NewSeckillPrePlaceOrderService(ctx context.Context) *SeckillPrePlaceOrderService {
	return &SeckillPrePlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *SeckillPrePlaceOrderService) Run(req *order.SeckillPrePlaceOrderReq) (resp *order.SeckillPrePlaceOrderResp, err error) {
	userId := req.UserId
	productId := req.ProductId
	if userId == 0 || productId == 0 {
		klog.CtxErrorf(s.ctx, "invalid params: userId=%d, productId=%d", userId, productId)
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "invalid params")
	}

	// todo rate limit

	seckillScript := script.GetPreSeckillScript()
	preOrderId, err := redis.NextId(s.ctx, "pre_order")
	if err != nil {
		klog.CtxErrorf(s.ctx, "redis.NextId.err:%v", err)
		return nil, kerrors.NewBizStatusError(errno.ErrInternal, "redis.NextId error")
	}
	productStockKey := redis.GetProductStockKey(productId)
	productOrderKey := redis.GetProductOrderKey(productId)
	preOrderKey := redis.GetOrderPreOrderKey(preOrderId)

	result, err := redis.RedisClient.Eval(s.ctx, seckillScript, []string{productStockKey, productOrderKey, preOrderKey}, userId, productId).Result()
	if err != nil {
		klog.CtxErrorf(s.ctx, "redis.RedisClient.Eval.err:%v", err)
		return nil, kerrors.NewBizStatusError(errno.ErrInternal, "redis.RedisClient.Eval error")
	}

	if resultMap, ok := result.(map[string]interface{}); ok {
		if errMsg, exists := resultMap["err"]; exists {
			switch errMsg {
			case "OUT_OF_STOCK":
				return nil, kerrors.NewBizStatusError(consts.ErrOutStock, "out of stock")
			case "DUPLICATE_USER":
				return nil, kerrors.NewBizStatusError(consts.ErrDuplicateUser, "duplicate user")
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
		klog.CtxErrorf(s.ctx, "producer.PublishPreOrder.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrPubMessage, "publish pre order message error")
	}

	resp = &order.SeckillPrePlaceOrderResp{
		TempId: preOrderId,
	}
	return
}
