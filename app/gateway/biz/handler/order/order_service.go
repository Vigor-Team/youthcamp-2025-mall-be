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

package order

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/order"
	"github.com/cloudwego/hertz/pkg/app"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// SeckillPrePlaceOrder .
// @router /order/seckill/pre [POST]
func SeckillPrePlaceOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.SeckillPrePlaceOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	resp, err := service.NewSeckillPrePlaceOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// SeckillPlaceOrder .
// @router /order/seckill [POST]
func SeckillPlaceOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.SeckillPlaceOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewSeckillPlaceOrderService(ctx, c).Run(&req)

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}

// QueryOrder .
// @router /order/:order_id [GET]
func QueryOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.QueryOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp, err := service.NewQueryOrderService(ctx, c).Run(&req)

	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}
	utils.SuccessResponse(c, resp)
}
