package auth

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreatePermission .
// @router /api/v1/permission/create [POST]
func CreatePermission(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.CreatePermissionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &auth.CreatePermissionResp{}
	resp, err = service.NewCreatePermissionService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// BindPermissionRole .
// @router /api/v1/permission/bind [POST]
func BindPermissionRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.BindPermissionRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &auth.BindPermissionRoleResp{}
	resp, err = service.NewBindPermissionRoleService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}
