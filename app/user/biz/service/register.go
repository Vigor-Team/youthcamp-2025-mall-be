// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Password != req.ConfirmPassword {
		err = kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "password and confirm password are not the same")
		return
	}
	userInfo, _ := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if userInfo != nil && userInfo.ID != 0 {
		err = kerrors.NewBizStatusError(consts.ErrUserExist, "email already exists")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		klog.CtxErrorf(s.ctx, "bcrypt.GenerateFromPassword: %v", err)
		err = kerrors.NewBizStatusError(errno.ErrInternal, "bcrypt.GenerateFromPassword error")
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		klog.CtxErrorf(s.ctx, "model.Create: %v", err)
		err = kerrors.NewBizStatusError(consts.ErrCreateUser, "create user error")
		return
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
