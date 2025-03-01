package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type AboutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAboutService(Context context.Context, RequestContext *app.RequestContext) *AboutService {
	return &AboutService{RequestContext: RequestContext, Context: Context}
}

func (h *AboutService) Run(req *common.Empty) (resp map[string]any, err error) {
	return utils.H{
		"title": "About",
	}, nil
}
