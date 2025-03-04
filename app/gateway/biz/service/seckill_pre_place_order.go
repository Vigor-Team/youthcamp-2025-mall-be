package service

import (
	"context"

	order "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
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
	preOrderId, err := rpc.OrderClient.SeckillPrePlaceOrder(h.Context, &rpcorder.SeckillPrePlaceOrderReq{
		UserId:    gatewayutils.GetUserIdFromCtx(h.RequestContext),
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
