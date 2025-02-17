package service

import (
	"context"
	"errors"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
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
	tempMeta, err := validateTempId(s.ctx, tempId, req.UserId)
	if err != nil {
		return nil, err
	}

	productId, err := strconv.ParseUint(tempMeta["product_id"], 10, 32)
	if err != nil {
		return nil, err
	}
	orderId, err := redis.NextId(s.ctx, redis.OrderNode)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	resp = &order.SeckillPlaceOrderResp{
		Status:  "processing",
		OrderId: strconv.Itoa(int(orderId)),
	}
	return
}

func validateTempId(ctx context.Context, tempId, userId uint32) (map[string]string, error) {
	if tempId == 0 || userId == 0 {
		return nil, errors.New("invalid param")
	}
	// get tempId info from redis
	productOrderKey := redis.GetOrderPreOrderKey(tempId)
	tempIdInfo, err := redis.RedisClient.HGetAll(ctx, productOrderKey).Result()
	if err != nil || len(tempIdInfo) == 0 {
		return nil, errors.New("invalid tempId")
	}
	if tempIdInfo["user_id"] != strconv.Itoa(int(userId)) {
		return nil, errors.New("invalid tempId")
	}
	return tempIdInfo, nil
}
