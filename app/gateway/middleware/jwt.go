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
	accessTokenExpire  = time.Hour
	refreshTokenExpire = time.Hour * 24 * 7
	once               sync.Once
	jwtMd              *jwt.HertzJWTMiddleware
)

func GetJwtMd() *jwt.HertzJWTMiddleware {
	once.Do(func() {
		jwtMd, _ = initJwtMd()
		_ = jwtMd.MiddlewareInit()
	})
	return jwtMd
}

func initJwtMd() (middleware *jwt.HertzJWTMiddleware, err error) {
	middleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("youthcamp2025mallbe"),
		Timeout:     accessTokenExpire,
		MaxRefresh:  refreshTokenExpire,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if userId, ok := data.(int32); ok {
				return jwt.MapClaims{
					IdentityKey: userId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userId := int32(claims[IdentityKey].(float64))
			return userId
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
				return nil, err
			}
			return res.UserId, nil
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			utils.ErrorResponse(c, int32(code), message)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			utils.SuccessResponse(c, &types.Token{Token: token})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return
}
