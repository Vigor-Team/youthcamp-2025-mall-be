package auth

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateRole .
// @router /api/v1/role/create [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.CreateRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &auth.CreateRoleResp{}
	resp, err = service.NewCreateRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// BindRoleUser .
// @router /api/v1/role/bind [POST]
func BindRoleUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.BindRoleUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &auth.BindRoleUserResp{}
	resp, err = service.NewBindRoleUserService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}
