package middleware

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	rpcuser "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"sync"
	"time"
)

var (
	IdentityKey        = "id"
	RefreshKey         = "refresh"
	accessTokenExpire  = time.Hour
	refreshTokenExpire = time.Hour * 24 * 7
	once               sync.Once
	jwtMd              *jwt.HertzJWTMiddleware
)

func GetJwtMd() *jwt.HertzJWTMiddleware {
	once.Do(func() {
		jwtMd, _ = initJwtMd()
	})
	return jwtMd
}

func initJwtMd() (middleware *jwt.HertzJWTMiddleware, err error) {
	middleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     accessTokenExpire,
		MaxRefresh:  refreshTokenExpire,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*types.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &types.User{
				ID: claims[IdentityKey].(uint32),
			}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals auth.LoginReq
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(loginVals.Email) == 0 || len(loginVals.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			res, err := rpc.UserClient.Login(ctx, &rpcuser.LoginReq{Email: loginVals.Email, Password: loginVals.Password})
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return res.UserId, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			// todo check permission
			return true
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			utils.ErrorResponse(c, code, message)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			utils.SuccessResponse(c, code, token)
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			utils.SuccessResponse(c, code, nil)
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return
}
