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
	tempId := req.TempId
	tempMeta, err := validateTempId(s.ctx, tempId)
	productId, err := strconv.ParseUint(tempMeta["product_id"], 10, 32)
	if err != nil {
		return nil, err
	}
	orderId, err := redis.NextId(s.ctx, "order_id")
	if err != nil {
		return nil, err
	}
	producer := mq.NewProducer(mq.Client)
	msg := mq.OrderMessage{
		TempID:       tempId,
		OrderId:      orderId,
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

func validateTempId(ctx context.Context, tempId string) (map[string]string, error) {
	if tempId == "" {
		return nil, errors.New("invalid tempId")
	}
	tempIdKey := redis.GetSeckillTempKey(tempId)
	tempIdInfo, err := redis.RedisClient.HGetAll(ctx, tempIdKey).Result()
	if err != nil || len(tempIdInfo) == 0 {
		return nil, errors.New("invalid tempId")
	}
	if tempIdInfo["status"] != "pre_held" {
		return nil, errors.New("status error")
	}
	redis.RedisClient.Del(ctx, tempIdKey)
	return tempIdInfo, nil
}
