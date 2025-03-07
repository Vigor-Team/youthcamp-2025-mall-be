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

// Code generated by hertz generator.

package auth

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	return nil
}

func _logoutMw() []app.HandlerFunc {
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _meMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _refreshMw() []app.HandlerFunc {
	return nil
}

func _bindpermissionroleMw() []app.HandlerFunc {
	return nil
}

func _createpermissionMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _roleMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _bindroleuserMw() []app.HandlerFunc {
	return nil
}

func _createroleMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _banuserMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _listpermissionsMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _listrolesMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _rolesMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _permissionsMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _updatepermissionMw() []app.HandlerFunc {
	var mws []app.HandlerFunc
	mws = append(mws, middleware.GetJwtMd().MiddlewareFunc())
	mws = append(mws, middleware.BlacklistMiddleware())
	mws = append(mws, middleware.CasbinAuth())
	return mws
}

func _unbindpermissionroleMw() []app.HandlerFunc {
	// your code...
	return nil
}
