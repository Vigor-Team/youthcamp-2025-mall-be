package service

import (
	"context"

	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchReq) (resp *product.ListProductsResp, err error) {
	r, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Query,
	})
	if err != nil {
		return
	}
	products := make([]*product.Product, 0, len(r.Results))
	for _, p := range r.Results {
		categoryNames := make([]string, 0, len(p.Categories))
		for _, c := range p.Categories {
			categoryNames = append(categoryNames, c)
		}
		products = append(products, &product.Product{
			ProductId:   p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Picture:     p.Picture,
			SpuPrice:    p.SpuPrice,
			SpuName:     p.SpuName,
			Stock:       p.Stock,
			Categories:  categoryNames,
		})
	}
	resp = &product.ListProductsResp{
		Products: products,
	}
	return
}
