package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

type SeckillPlaceOrderService struct {
	ctx context.Context
} // NewSeckillPlaceOrderService new SeckillPlaceOrderService
func NewSeckillPlaceOrderService(ctx context.Context) *SeckillPlaceOrderService {
	return &SeckillPlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *SeckillPlaceOrderService) Run(req *order.SeckillPlaceOrderReq) (resp *order.SeckillPlaceOrderResp, err error) {
	// validate tempId
	tempId := req.TempId
	tempMeta, err := validatePreOrderId(s.ctx, tempId, req.UserId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "validatePreOrderId.err:%v", err)
		return nil, consts.ErrInvalidParams
	}

	productId, err := strconv.ParseUint(tempMeta["product_id"], 10, 32)
	if err != nil {
		klog.CtxErrorf(s.ctx, "strconv.ParseUint.err:%v", err)
		return nil, consts.ErrInvalidParams
	}
	orderId, err := redis.NextId(s.ctx, "order")
	if err != nil {
		klog.CtxErrorf(s.ctx, "redis.NextId.err:%v", err)
		return nil, consts.ErrRedis
	}

	// publish order message
	producer := mq.NewProducer(mq.Client)
	msg := mq.OrderMessage{
		TempID:       strconv.Itoa(int(tempId)),
		OrderId:      orderId,
		UserID:       req.UserId,
		UserCurrency: req.UserCurrency,
		ProductId:    uint32(productId),
		Consignee: model.Consignee{
			Email:         req.Email,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	}
	err = producer.PublishOrder(s.ctx, msg)
	if err != nil {
		klog.CtxErrorf(s.ctx, "producer.PublishOrder.err:%v", err)
		return nil, consts.ErrPublishMessage
	}
	resp = &order.SeckillPlaceOrderResp{
		Status:  "processing",
		OrderId: strconv.Itoa(int(orderId)),
	}
	return
}

func validatePreOrderId(ctx context.Context, tempId, userId uint32) (map[string]string, error) {
	if tempId == 0 || userId == 0 {
		return nil, fmt.Errorf("invalid params: tempId=%d, userId=%d", tempId, userId)
	}
	// get tempId info from redis
	productOrderKey := redis.GetOrderPreOrderKey(tempId)
	tempIdInfo, err := redis.RedisClient.HGetAll(ctx, productOrderKey).Result()
	fmt.Println("tempIdInfo", tempIdInfo)
	if err != nil || len(tempIdInfo) == 0 {
		return nil, fmt.Errorf("invalid tempId: %v", tempId)
	}
	if tempIdInfo["user_id"] != strconv.Itoa(int(userId)) || tempIdInfo["product_id"] == "" {
		return nil, fmt.Errorf("invalid tempId: %v", tempId)
	}
	return tempIdInfo, nil
}
