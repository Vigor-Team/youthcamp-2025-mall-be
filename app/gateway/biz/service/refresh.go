package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/_consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	"github.com/hertz-contrib/jwt"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
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
		return nil, _consts.ErrTokenNotFound
	}
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, _consts.ErrTokenFormat
	}
	token, err := h.jwtMd.ParseTokenString(parts[1])
	if err != nil {
		return nil, _consts.ErrTokenInvalid
	}
	claims := jwt.ExtractClaimsFromToken(token)
	if claims == nil {
		return nil, _consts.ErrTokenInvalid
	}
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp+int64(h.jwtMd.MaxRefresh.Seconds()) {
		return nil, _consts.ErrTokenExpired
	}
	newToken, _, err := h.jwtMd.TokenGenerator(claims["id"])
	if err != nil {
		return
	}
	return &types.Token{
		Token: newToken,
	}, nil
}
