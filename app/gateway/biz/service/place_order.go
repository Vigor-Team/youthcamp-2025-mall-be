package service

import (
	"context"

	order "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	r, err := rpc.OrderClient.PlaceOrder(h.Context, &rpcorder.PlaceOrderReq{
		UserId:       gatewayutils.GetUserIdFromCtx(h.RequestContext),
		UserCurrency: req.UserCurrency,
		Address: &rpcorder.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Email: req.Email,
		OrderItems: func() (ois []*rpcorder.OrderItem) {
			ois = make([]*rpcorder.OrderItem, 0, len(req.OrderItems))
			for _, oi := range req.OrderItems {
				ois = append(ois, &rpcorder.OrderItem{
					Item: &rpccart.CartItem{
						ProductId: oi.Item.ProductId,
						Quantity:  oi.Item.Quantity,
					},
					Cost: oi.Cost,
				})
			}
			return ois
		}(),
	})
	if err != nil {
		return
	}
	resp = &order.PlaceOrderResp{
		OrderId: r.Order.OrderId,
	}
	return
}
