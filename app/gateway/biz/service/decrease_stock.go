package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DecreaseStockService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDecreaseStockService(Context context.Context, RequestContext *app.RequestContext) *DecreaseStockService {
	return &DecreaseStockService{RequestContext: RequestContext, Context: Context}
}

func (h *DecreaseStockService) Run(req *product.StockOpReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.DecrStock(h.Context, &rpcproduct.DecrStockReq{
		Id:   req.ProductId,
		Decr: req.Quantity,
	})
	if err != nil {
		return
	}
	return
}
