package service

import (
	"context"

	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type DeleteCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCartItemService(Context context.Context, RequestContext *app.RequestContext) *DeleteCartItemService {
	return &DeleteCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCartItemService) Run(req *cart.DeleteCartItemReq) (resp *cart.DeleteCartItemResp, err error) {
	_, err = rpc.CartClient.DeleteItem(
		h.Context, &rpccart.DeleteItemReq{
			UserId:    gatewayutils.GetUserIdFromCtx(h.RequestContext),
			ProductId: req.ProductId,
		},
	)
	if err != nil {
		hlog.CtxErrorf(h.Context, "DeleteCartItemService Run failed: %v", err)
		return nil, err
	}
	return
}
