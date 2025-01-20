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
	"net/http"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"

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

func ErrorResponse(c *app.RequestContext, code int32, message string) {
	c.JSON(http.StatusOK, GlobalResponse{
		Code: code,
		Msg:  message,
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

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	var cartNum int
	userId := gatewayutils.GetUserIdFromCtx(ctx)
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: userId})
	if cartResp != nil && cartResp.Cart != nil {
		cartNum = len(cartResp.Cart.Items)
	}
	content["user_id"] = ctx.Value(gatewayutils.UserIdKey)
	content["cart_num"] = cartNum
	return content
}
