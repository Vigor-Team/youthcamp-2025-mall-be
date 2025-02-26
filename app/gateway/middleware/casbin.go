package middleware

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/conf"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadpter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"os"
)

var AuthEnforcer *casbin.Enforcer

func InitCasbin() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	adapter, err := gormadpter.NewAdapter("mysql", dsn, true)
	if err != nil {
		panic(err)
	}

	m, err := model.NewModelFromString(
		`
#request input
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`,
	)

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	AuthEnforcer = enforcer
}

func Authorize(rvals ...interface{}) (result bool, err error) {
	// casbin enforce
	res, err1 := AuthEnforcer.Enforce(rvals[0], rvals[1], rvals[2])
	return res, err1
}

func CasbinAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		roles, exists := c.Get(Roles)
		fmt.Println("roles: ", roles)
		if !exists {
			utils.FailResponse(ctx, c, kerrors.NewBizStatusError(1, "role not found"))
			c.Abort()
			return
		}

		// casbin enforce
		var isAuth = false

		for _, r := range roles.([]interface{}) {
			fmt.Println("role: ", r)
			fmt.Println("path: ", c.FullPath())
			fmt.Println("method: ", string(c.Request.Header.Method()))
			res, err := Authorize(r.(string), c.FullPath(), string(c.Request.Header.Method()))
			if err != nil {
				hlog.CtxErrorf(ctx, "Authorize is error: %v", err)
				utils.FailResponse(ctx, c, kerrors.NewBizStatusError(1, "Authorize is error"))
				c.Abort()
				return
			}
			if res {
				isAuth = true
				break
			}
		}

		if isAuth {
			c.Next(ctx)
		} else {
			utils.FailResponse(ctx, c, kerrors.NewBizStatusError(1, "FORBIDDEN"))
			c.Abort()
			return
		}
	}
}
