package middleware

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcuser "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hertz-contrib/jwt"
	"sync"
	"time"
)

var (
	UserID = "user_id"
	Roles  = "roles"

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
		IdentityKey: UserID,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if claim, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					UserID: claim[UserID],
					Roles:  claim[Roles],
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userID := int32(claims[UserID].(float64))
			roles := claims[Roles]
			c.Set(Roles, roles)
			return userID
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals auth.LoginReq
			if err = c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(loginVals.Email) == 0 || len(loginVals.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			loginRes, err := rpc.UserClient.Login(ctx, &rpcuser.LoginReq{Email: loginVals.Email, Password: loginVals.Password})
			if err != nil {
				return nil, err
			}
			userID := loginRes.UserId
			roles, err := model.GetRolesByUid(mysql.DB, ctx, int64(userID))
			if err != nil {
				return nil, kerrors.NewBizStatusError(10001, "get user roles error")
			}
			roleNames := make([]string, 0, len(roles))
			for _, role := range roles {
				roleNames = append(roleNames, role.Name)
			}
			return map[string]interface{}{
				UserID: loginRes.UserId,
				Roles:  roleNames,
			}, nil
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			utils.ErrorResponse(c, int32(code), message)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return
}
