package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	jwtMd          *jwt.HertzJWTMiddleware
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext, jwtMd *jwt.HertzJWTMiddleware) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context, jwtMd: jwtMd}
}

func (h *LoginService) Run(_ *auth.LoginReq) (resp *types.Token, err error) {
	authRes, err := h.jwtMd.Authenticator(h.Context, h.RequestContext)
	if err != nil {
		hlog.CtxErrorf(h.Context, "login error: %v", err)
		return nil, err
	}
	token, _, err := h.jwtMd.TokenGenerator(authRes)
	if err != nil {
		hlog.CtxErrorf(h.Context, "accessToken gen error: %v", err)
		return nil, err
	}
	resp = &types.Token{
		Token: token,
	}
	return
}
