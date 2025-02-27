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

package utils

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	defaultSuccessCode = 0
	defaultSuccessMsg  = "success"
)

type GlobalResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func FailResponse(ctx context.Context, c *app.RequestContext, err error) {
	if kerr, ok := kerrors.FromBizStatusError(err); ok {
		hlog.CtxErrorf(ctx, "biz error: %v", kerr)
		FailResponseWithCodeAndMsg(c, kerr.BizStatusCode(), kerr.BizMessage())
		return
	}
	hlog.CtxErrorf(ctx, "unknown error: %v", err)
	FailResponseWithMsg(c, "unknown error, please try again later")
}

func FailResponseWithMsg(c *app.RequestContext, msg string) {
	c.JSON(http.StatusOK, GlobalResponse{
		Code: errno.ErrInternal,
		Msg:  msg,
		Data: nil,
	})
}

func FailResponseWithCodeAndMsg(c *app.RequestContext, code int32, msg string) {
	c.JSON(http.StatusOK, GlobalResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func SuccessResponse(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, GlobalResponse{
		Code: defaultSuccessCode,
		Msg:  defaultSuccessMsg,
		Data: data,
	})
}
