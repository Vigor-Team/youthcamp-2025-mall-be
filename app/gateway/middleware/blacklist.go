package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// BlacklistMiddleware is a middleware that checks if the request is blacklisted.
func BlacklistMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 从JWT Claims中获取角色

		// 从redis检查userId

		// 从redis检查ip

		// 从redis检查token

		// 放行
	}
}
