// Code generated by hertz generator. DO NOT EDIT.

package auth

import (
	auth "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/handler/auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			_v1.POST("/ban", append(_banuserMw(), auth.BanUser)...)
			_v1.POST("/login", append(_loginMw(), auth.Login)...)
			_v1.POST("/logout", append(_logoutMw(), auth.Logout)...)
			_v1.GET("/me", append(_meMw(), auth.Me)...)
			_v1.POST("/permissions", append(_createpermissionMw(), auth.CreatePermission)...)
			_v1.PUT("/permissions", append(_updatepermissionMw(), auth.UpdatePermission)...)
			_v1.GET("/refresh", append(_refreshMw(), auth.Refresh)...)
			_v1.POST("/register", append(_registerMw(), auth.Register)...)
			_v1.GET("/roles", append(_listrolesMw(), auth.ListRoles)...)
			_roles := _v1.Group("/roles", _rolesMw()...)
			_roles.POST("/bind", append(_bindroleuserMw(), auth.BindRoleUser)...)
			_v1.POST("/roles", append(_createroleMw(), auth.CreateRole)...)
			_v1.GET("/permissions", append(_listpermissionsMw(), auth.ListPermissions)...)
			_permissions := _v1.Group("/permissions", _permissionsMw()...)
			_permissions.POST("/bind", append(_bindpermissionroleMw(), auth.BindPermissionRole)...)
			_permissions.POST("/unbind", append(_unbindpermissionroleMw(), auth.UnbindPermissionRole)...)
		}
	}
}
