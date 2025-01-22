package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/model"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
)

type UserInfoService struct {
	ctx context.Context
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	//userId := req.UserId
	userRow, err := model.GetByID(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return
	}
	return &user.UserInfoResp{Email: userRow.Email, Role: userRow.Role}, nil
}
