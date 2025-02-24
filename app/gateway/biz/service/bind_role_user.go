package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type BindRoleUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBindRoleUserService(Context context.Context, RequestContext *app.RequestContext) *BindRoleUserService {
	return &BindRoleUserService{RequestContext: RequestContext, Context: Context}
}

func (h *BindRoleUserService) Run(req *auth.BindRoleUserReq) (resp *auth.BindRoleUserResp, err error) {
	err = model.BindUserRole(mysql.DB, h.Context, &model.UserRole{
		UID: req.Uid,
		RID: req.Rid,
	})
	if err != nil {
		return nil, err
	}
	return
}
