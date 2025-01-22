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
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/types"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"

	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	jwtMd          *jwt.HertzJWTMiddleware
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext, jwtMd *jwt.HertzJWTMiddleware) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context, jwtMd: jwtMd}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *types.Token, err error) {
	authRes, err := h.jwtMd.Authenticator(h.Context, h.RequestContext)
	if err != nil {
		hlog.CtxErrorf(h.Context, "login error: %v", err)
		return nil, err
	}
	userId := authRes.(int32)
	token, _, err := h.jwtMd.TokenGenerator(userId)
	if err != nil {
		hlog.CtxErrorf(h.Context, "accessToken gen error: %v", err)
		return nil, err
	}
	resp = &types.Token{
		Token: token,
	}
	return
}
