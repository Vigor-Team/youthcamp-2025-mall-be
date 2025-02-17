package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	order "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type SeckillPrePlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSeckillPrePlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *SeckillPrePlaceOrderService {
	return &SeckillPrePlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *SeckillPrePlaceOrderService) Run(req *order.SeckillPrePlaceOrderReq) (resp *order.SeckillPrePlaceOrderResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	preOrderId, err := rpc.OrderClient.SeckillPrePlaceOrder(h.Context, &rpcorder.SeckillPrePlaceOrderReq{
		UserId:    1,
		ProductId: req.ProductId,
	})
	if err != nil {
		return
	}
	resp = &order.SeckillPrePlaceOrderResp{
		TempId: preOrderId.TempId,
	}
	return
}
