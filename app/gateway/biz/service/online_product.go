package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type OnlineProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOnlineProductService(Context context.Context, RequestContext *app.RequestContext) *OnlineProductService {
	return &OnlineProductService{RequestContext: RequestContext, Context: Context}
}

func (h *OnlineProductService) Run(req *product.ProductIDReq) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	_, err = rpc.ProductClient.OnlineProduct(h.Context, &rpcproduct.OnlineProductReq{
		Id: req.ProductId,
	})
	if err != nil {
		return
	}
	return
}
