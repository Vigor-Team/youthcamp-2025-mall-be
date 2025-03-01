package service

import (
	"context"

	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCartItemService(Context context.Context, RequestContext *app.RequestContext) *UpdateCartItemService {
	return &UpdateCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCartItemService) Run(req *cart.UpdateCartReq) (resp *cart.UpdateCartResp, err error) {
	_, err = rpc.CartClient.UpdateCart(
		h.Context, &rpccart.UpdateCartReq{
			UserId: uint32(h.RequestContext.Value("user_id").(int32)),
			Item: &rpccart.CartItem{
				ProductId: req.ProductId,
				Quantity:  req.ProductNum,
			},
		},
	)
	return
}
