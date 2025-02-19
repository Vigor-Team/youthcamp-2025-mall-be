package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type BindPermissionRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBindPermissionRoleService(Context context.Context, RequestContext *app.RequestContext) *BindPermissionRoleService {
	return &BindPermissionRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *BindPermissionRoleService) Run(req *auth.BindPermissionRoleReq) (resp *auth.BindPermissionRoleResp, err error) {
	err = model.BindPermissionRole(mysql.DB, h.Context, &model.PermissionRole{
		PID: req.Pid,
		RID: req.Rid,
	})
	resp = &auth.BindPermissionRoleResp{
		PermissionRole: &auth.PermissionRole{
			Pid: req.Pid,
			Rid: req.Rid,
		},
	}
	return
}
