package middleware

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// BlacklistMiddleware is a middleware that checks if the request is blacklisted.
func BlacklistMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userID, exists := c.Get(UserID)
		if !exists {
			utils.FailResponse(ctx, c, kerrors.NewBizStatusError(1, "role not found"))
			c.Abort()
			return
		}
		userIDKey := redis.GetBlacklistUserIDKey(userID.(int64))
		if redis.RedisClient.Exists(ctx, userIDKey).Val() == 1 {
			utils.FailResponse(ctx, c, kerrors.NewBizStatusError(1, "user is blacklisted"))
			c.Abort()
		}
		c.Next(ctx)
	}
}
