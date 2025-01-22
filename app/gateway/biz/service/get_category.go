package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
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

func (h *GetCategoryService) Run(req *product.GetCategoryReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
