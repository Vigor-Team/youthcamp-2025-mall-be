package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type IncreaseStockService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewIncreaseStockService(Context context.Context, RequestContext *app.RequestContext) *IncreaseStockService {
	return &IncreaseStockService{RequestContext: RequestContext, Context: Context}
}

func (h *IncreaseStockService) Run(req *product.StockOpReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.IncrStock(h.Context, &rpcproduct.IncrStockReq{
		Id:   req.ProductId,
		Incr: req.Quantity,
	})
	if err != nil {
		return
	}
	return
}
