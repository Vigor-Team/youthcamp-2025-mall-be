package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateRoleService(Context context.Context, RequestContext *app.RequestContext) *CreateRoleService {
	return &CreateRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateRoleService) Run(req *auth.CreateRoleReq) (resp *auth.CreateRoleResp, err error) {
	role, err := model.CreateRole(mysql.DB, h.Context, &model.Role{
		Name: req.Name,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "model.CreateRole.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrCreateRole, err.Error())
	}
	resp = &auth.CreateRoleResp{
		Role: &auth.Role{
			Name: role.Name,
		},
	}
	return
}
