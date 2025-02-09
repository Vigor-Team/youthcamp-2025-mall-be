package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryService {
	return &GetCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryService) Run(req *product.GetCategoryReq) (resp *product.GetCategoryResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	r, err := rpc.ProductClient.GetCategory(h.Context, &rpcproduct.GetCategoryReq{
		Id: req.CategoryId,
	})
	if err != nil {
		return
	}
	resp = &product.GetCategoryResp{
		Category: &product.Category{
			CategoryId:  r.Category.Id,
			Name:        r.Category.Name,
			Description: r.Category.Description,
		},
	}
	return
}
