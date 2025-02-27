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
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/model"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.GetByEmail: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUserNotFound, "user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		klog.CtxErrorf(s.ctx, "bcrypt.CompareHashAndPassword: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrPassword, "password error")
	}
	return &user.LoginResp{UserId: int32(userRow.ID)}, nil
}
