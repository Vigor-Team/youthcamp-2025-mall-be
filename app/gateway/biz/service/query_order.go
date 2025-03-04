package service

import (
	"context"

	order "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type QueryOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryOrderService(Context context.Context, RequestContext *app.RequestContext) *QueryOrderService {
	return &QueryOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryOrderService) Run(req *order.QueryOrderReq) (resp *order.QueryOrderResp, err error) {
	res, err := rpc.OrderClient.QueryOrder(h.Context, &rpcorder.QueryOrderReq{
		OrderId: req.OrderId,
		UserId:  gatewayutils.GetUserIdFromCtx(h.RequestContext),
	})
	if err != nil {
		return
	}
	var ois []*order.OrderItem
	for _, oi := range res.Order.OrderItems {
		ois = append(ois, &order.OrderItem{
			Item: &order.CartItem{
				ProductId: oi.Item.ProductId,
				Quantity:  oi.Item.Quantity,
			},
			Cost: oi.Cost,
		})
	}
	resp = &order.QueryOrderResp{
		Order: &order.Order{
			OrderId:      res.Order.OrderId,
			OrderItems:   ois,
			UserId:       res.Order.UserId,
			UserCurrency: res.Order.UserCurrency,
			Address: &order.Address{
				StreetAddress: res.Order.Address.StreetAddress,
				City:          res.Order.Address.City,
				State:         res.Order.Address.State,
				Country:       res.Order.Address.Country,
				ZipCode:       res.Order.Address.ZipCode,
			},
			Email:     res.Order.Email,
			Status:    res.Order.Status,
			CreatedAt: res.Order.CreatedAt,
		},
	}
	return
}
