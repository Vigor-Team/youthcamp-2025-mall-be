package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/redis"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type BanUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBanUserService(Context context.Context, RequestContext *app.RequestContext) *BanUserService {
	return &BanUserService{RequestContext: RequestContext, Context: Context}
}

func (h *BanUserService) Run(req *auth.BanUserReq) (resp *auth.BanUserResp, err error) {
	userIDKey := redis.GetBlacklistUserIDKey(int32(req.UserId))
	err = redis.RedisClient.Set(h.Context, userIDKey, 1, time.Duration(req.ExpireSeconds)*time.Second).Err()
	if err != nil {
		klog.CtxErrorf(h.Context, "RedisClient.Set.err: %v", err)
		err = kerrors.NewBizStatusError(consts.ErrRedisErr, err.Error())
	}
	return
}
