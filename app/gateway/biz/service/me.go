package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcuser "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type MeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMeService(Context context.Context, RequestContext *app.RequestContext) *MeService {
	return &MeService{RequestContext: RequestContext, Context: Context}
}

func (h *MeService) Run() (resp *auth.MeResp, err error) {
	userId := gatewayutils.GetUserIdFromCtx(h.RequestContext)
	res, err := rpc.UserClient.UserInfo(h.Context, &rpcuser.UserInfoReq{UserId: int32(userId)})
	if err != nil {
		return nil, err
	}
	role, err := model.GetRolesByUid(mysql.DB, h.Context, int64(userId))
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetRolesByUid error: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetRole, "GetRolesByUid error")
	}
	rs := make([]*auth.Role, 0, len(role))
	for _, r := range role {
		rs = append(rs, &auth.Role{
			Id:   r.ID,
			Name: r.Name,
		})
	}
	return &auth.MeResp{
		Id:    int64(userId),
		Email: res.Email,
		Roles: rs,
	}, err
}
