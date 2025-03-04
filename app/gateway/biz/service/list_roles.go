package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListRolesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListRolesService(Context context.Context, RequestContext *app.RequestContext) *ListRolesService {
	return &ListRolesService{RequestContext: RequestContext, Context: Context}
}

func (h *ListRolesService) Run(_ *common.Empty) (resp *auth.ListRolesResp, err error) {
	roles, err := model.ListRoles(mysql.DB, h.Context)
	if err != nil {
		klog.CtxErrorf(h.Context, "model.ListRoles.err: %v", err)
		err = kerrors.NewBizStatusError(consts.ErrGetRole, "get role failed")
	}
	rs := make([]*auth.Role, 0, len(roles))
	for _, r := range roles {
		rs = append(rs, &auth.Role{
			Id:   r.ID,
			Name: r.Name,
		})
	}
	resp = &auth.ListRolesResp{
		Roles: rs,
	}
	return
}
