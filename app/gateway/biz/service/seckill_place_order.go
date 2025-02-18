package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	order "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type SeckillPlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSeckillPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *SeckillPlaceOrderService {
	return &SeckillPlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *SeckillPlaceOrderService) Run(req *order.SeckillPlaceOrderReq) (resp *order.SeckillPlaceOrderResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	seckill, err := rpc.OrderClient.SeckillPlaceOrder(h.Context, &rpcorder.SeckillPlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address: &rpcorder.Address{
			StreetAddress: req.Address.StreetAddress,
			ZipCode:       req.Address.ZipCode,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
		},
		Email:  req.Email,
		TempId: req.TempId,
	})
	if err != nil {
		return
	}
	resp = &order.SeckillPlaceOrderResp{
		OrderId: seckill.OrderId,
		Status:  seckill.Status,
	}
	return
}
