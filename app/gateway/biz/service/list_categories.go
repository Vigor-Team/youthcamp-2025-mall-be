package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListCategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoriesService(Context context.Context, RequestContext *app.RequestContext) *ListCategoriesService {
	return &ListCategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoriesService) Run(req *common.Empty) (resp *product.ListCategoriesResp, err error) {
	r, err := rpc.ProductClient.GetCategories(h.Context, &rpcproduct.GetCategoriesReq{})
	if err != nil {
		return
	}
	categories := make([]*product.Category, 0, len(r.Categories))
	for _, c := range r.Categories {
		categories = append(categories, &product.Category{
			CategoryId:  c.Id,
			Name:        c.Name,
			Description: c.Description,
		})
	}
	resp = &product.ListCategoriesResp{
		Categories: categories,
	}
	return
}
