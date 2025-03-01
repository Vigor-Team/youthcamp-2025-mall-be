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

type ListPermissionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListPermissionsService(Context context.Context, RequestContext *app.RequestContext) *ListPermissionsService {
	return &ListPermissionsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListPermissionsService) Run(_ *common.Empty) (resp *auth.ListPermissionsResp, err error) {
	permissions, err := model.ListPermissions(mysql.DB, h.Context)
	if err != nil {
		klog.CtxErrorf(h.Context, "model.ListPermissions.err: %v", err)
		err = kerrors.NewBizStatusError(consts.ErrGetPermission, "get permission failed")
	}
	ps := make([]*auth.Permission, 0, len(permissions))
	for _, p := range permissions {
		ps = append(ps, &auth.Permission{
			Id: p.ID,
			V1: p.V1,
			V2: p.V2,
		})
	}
	resp = &auth.ListPermissionsResp{
		Permissions: ps,
	}
	return
}
