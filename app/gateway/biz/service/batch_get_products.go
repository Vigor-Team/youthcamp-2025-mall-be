package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type BatchGetProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBatchGetProductsService(Context context.Context, RequestContext *app.RequestContext) *BatchGetProductsService {
	return &BatchGetProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *BatchGetProductsService) Run(req *product.BatchProductsReq) (resp *product.BatchProductsResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	r, err := rpc.ProductClient.BatchGetProducts(h.Context, &rpcproduct.BatchGetProductsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return
	}

	products := map[uint32]*product.Product{}
	for _, p := range r.Products {
		categoryNames := make([]string, 0, len(p.Categories))
		for _, c := range p.Categories {
			categoryNames = append(categoryNames, c)
		}
		products[p.Id] = &product.Product{
			ProductId:   p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Picture:     p.Picture,
			SpuPrice:    p.SpuPrice,
			SpuName:     p.SpuName,
			Stock:       p.Stock,
			Categories:  categoryNames,
		}
	}
	resp = &product.BatchProductsResp{
		Products: products,
	}
	return
}
