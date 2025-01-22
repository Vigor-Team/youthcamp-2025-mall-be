package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/middleware"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	rpcuser "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type MeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMeService(Context context.Context, RequestContext *app.RequestContext) *MeService {
	return &MeService{RequestContext: RequestContext, Context: Context}
}

func (h *MeService) Run() (resp *types.UserInfo, err error) {
	claims, err := middleware.GetJwtMd().GetClaimsFromJWT(h.Context, h.RequestContext)
	fmt.Println("claims", claims)
	if err != nil {
		return nil, err
	}
	userId, ok := claims["id"].(int32)
	fmt.Println("userId", userId)
	if !ok {
		return nil, err
	}
	res, err := rpc.UserClient.UserInfo(h.Context, &rpcuser.UserInfoReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Email: res.Email,
		Role:  res.Role,
	}, err
}
