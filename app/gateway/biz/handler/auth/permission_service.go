package auth

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/auth"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
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

// ListPermissions .
// @router /api/v1/permissions [GET]
func ListPermissions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewListPermissionsService(ctx, c).Run(&req)

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// UpdatePermission .
// @router /api/v1/permissions [PUT]
func UpdatePermission(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UpdatePermissionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewUpdatePermissionService(ctx, c).Run(&req)

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// UnbindPermissionRole .
// @router /api/v1/permissions/unbind [POST]
func UnbindPermissionRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UnbindPermissionRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewUnbindPermissionRoleService(ctx, c).Run(&req)

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}
