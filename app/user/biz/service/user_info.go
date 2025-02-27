package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UserInfoService struct {
	ctx context.Context
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	userRow, err := model.GetByID(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.GetByID: %v", err)
		err = kerrors.NewBizStatusError(errno.ErrMysql, "get user info error")
		return
	}
	if userRow == nil {
		err = kerrors.NewBizStatusError(consts.ErrUserNotFound, "user not found")
		return
	}

	return &user.UserInfoResp{Email: userRow.Email}, nil
}
