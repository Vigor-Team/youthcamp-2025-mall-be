package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type UpdatePermissionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePermissionService(Context context.Context, RequestContext *app.RequestContext) *UpdatePermissionService {
	return &UpdatePermissionService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePermissionService) Run(req *auth.UpdatePermissionReq) (resp *auth.UpdatePermissionResp, err error) {
	err = model.UpdatePermission(mysql.DB, h.Context, &model.Permission{
		ID: req.Id,
		V1: req.V1,
		V2: req.V2,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "UpdatePermission failed: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUpdatePermission, "UpdatePermission failed")
	}
	return
}
