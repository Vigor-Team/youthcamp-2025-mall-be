package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/model"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.UserDeleteReq) (resp *user.UserDeleteResp, err error) {
	userID := req.UserId
	u, err := model.GetByID(mysql.DB, s.ctx, uint(userID))
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.GetByID: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetUser, "get user error")
	}
	if u == nil {
		return nil, kerrors.NewBizStatusError(consts.ErrUserNotFound, "user not exist")
	}
	if err = model.Delete(mysql.DB, s.ctx, u); err != nil {
		klog.CtxErrorf(s.ctx, "model.Delete: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrDeleteUser, "delete user error")
	}
	return
}
