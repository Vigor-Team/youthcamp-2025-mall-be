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

type UnbindPermissionRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUnbindPermissionRoleService(Context context.Context, RequestContext *app.RequestContext) *UnbindPermissionRoleService {
	return &UnbindPermissionRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *UnbindPermissionRoleService) Run(req *auth.UnbindPermissionRoleReq) (resp *auth.UnbindPermissionRoleResp, err error) {
	err = model.UnbindPermissionRole(mysql.DB, h.Context, &model.PermissionRole{
		PID: req.Pid,
		RID: req.Rid,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "UnbindPermissionRole failed: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUnbindPermission, "UnbindPermissionRole failed")
	}
	resp = &auth.UnbindPermissionRoleResp{}
	return
}
