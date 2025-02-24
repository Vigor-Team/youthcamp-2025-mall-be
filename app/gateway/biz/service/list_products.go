package service

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsService(Context context.Context, RequestContext *app.RequestContext) *ListProductsService {
	return &ListProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	r, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{
		CategoryId: req.CategoryId,
		Role:       h.RequestContext.Value("roles").([]interface{})[0].(string),
	})
	if err != nil {
		return
	}
	products := make([]*product.Product, 0, len(r.Products))
	for _, p := range r.Products {
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
