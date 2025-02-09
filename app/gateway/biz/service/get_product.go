package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductIDReq) (resp *product.GetProductResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	r, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
		Id: req.ProductId,
	})
	if err != nil {
		return
	}
	resp = &product.GetProductResp{
		Product: &product.Product{
			ProductId:   r.Product.Id,
			Name:        r.Product.Name,
			Description: r.Product.Description,
			Price:       r.Product.Price,
			Picture:     r.Product.Picture,
			SpuPrice:    r.Product.SpuPrice,
			SpuName:     r.Product.SpuName,
			Stock:       r.Product.Stock,
			Categories:  r.Product.Categories,
		},
	}
	return
}
