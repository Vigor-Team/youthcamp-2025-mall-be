package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type OfflineProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOfflineProductService(Context context.Context, RequestContext *app.RequestContext) *OfflineProductService {
	return &OfflineProductService{RequestContext: RequestContext, Context: Context}
}

func (h *OfflineProductService) Run(req *product.ProductIDReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.OfflineProduct(h.Context, &rpcproduct.OfflineProductReq{
		Id: req.ProductId,
	})
	if err != nil {
		return
	}
	return
}
