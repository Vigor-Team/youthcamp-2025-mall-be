package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hertz-contrib/jwt"
)

type RefreshService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	jwtMd          *jwt.HertzJWTMiddleware
}

func NewRefreshService(Context context.Context, RequestContext *app.RequestContext, jwtMd *jwt.HertzJWTMiddleware) *RefreshService {
	return &RefreshService{RequestContext: RequestContext, Context: Context, jwtMd: jwtMd}
}

func (h *RefreshService) Run() (resp *types.Token, err error) {
	auth := h.RequestContext.Request.Header.Get("Authorization")
	if auth == "" {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTNotFound, "jwt not found")
	}
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTInvalid, "jwt invalid")
	}
	token, err := h.jwtMd.ParseTokenString(parts[1])
	if err != nil {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTInvalid, "jwt invalid")
	}
	claims := jwt.ExtractClaimsFromToken(token)
	if claims == nil {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTInvalid, "jwt invalid")
	}
	fmt.Println(claims)
	origIat := int64(claims["orig_iat"].(float64))
	if time.Now().Unix() > origIat+h.jwtMd.MaxRefresh.Milliseconds() {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTRefreshExpired, "jwt refresh expired")
	}
	newToken, _, err := h.jwtMd.TokenGenerator(claims)
	if err != nil {
		return nil, kerrors.NewBizStatusError(consts.ErrJWTCreate, "jwt create failed")
	}
	return &types.Token{
		Token: newToken,
	}, nil
}
