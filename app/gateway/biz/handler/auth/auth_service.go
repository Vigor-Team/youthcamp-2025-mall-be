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

package auth

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /api/v1/auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "sign-up", hertzUtils.H{"error": err})
		return
	}
}

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewLoginService(ctx, c, middleware.GetJwtMd()).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
}

// Me .
// @router /api/v1/me [GET]
func Me(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewMeService(ctx, c).Run()

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// Refresh .
// @router /api/v1/refresh [GET]
func Refresh(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewRefreshService(ctx, c, middleware.GetJwtMd()).Run()
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}
