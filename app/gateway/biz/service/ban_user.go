package service

import (
	"context"
	"time"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/dal/redis"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
		hlog.CtxErrorf(h.Context, "RedisClient.Set.err: %v", err)
		err = kerrors.NewBizStatusError(consts.ErrBanUser, err.Error())
		return
	}
	return
}
