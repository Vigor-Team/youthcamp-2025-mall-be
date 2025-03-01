package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreatePermissionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreatePermissionService(Context context.Context, RequestContext *app.RequestContext) *CreatePermissionService {
	return &CreatePermissionService{RequestContext: RequestContext, Context: Context}
}

func (h *CreatePermissionService) Run(req *auth.CreatePermissionReq) (resp *auth.CreatePermissionResp, err error) {
	permission, err := model.CreatePermission(mysql.DB, h.Context, &model.Permission{
		V1: req.V1,
		V2: req.V2,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "model.CreatePermission.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrCreatePermission, err.Error())
	}
	resp = &auth.CreatePermissionResp{
		Permission: &auth.Permission{
			V1: permission.V1,
			V2: permission.V2,
		},
	}
	return
}
