package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListCategoriesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListCategoriesService(Context context.Context, RequestContext *app.RequestContext) *ListCategoriesService {
	return &ListCategoriesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListCategoriesService) Run(req *common.Empty) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	return
}
