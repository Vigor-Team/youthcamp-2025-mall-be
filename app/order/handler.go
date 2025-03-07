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

package main

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/service"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// SeckillPlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SeckillPlaceOrder(ctx context.Context, req *order.SeckillPlaceOrderReq) (resp *order.SeckillPlaceOrderResp, err error) {
	resp, err = service.NewSeckillPlaceOrderService(ctx).Run(req)

	return resp, err
}

// SeckillPrePlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) SeckillPrePlaceOrder(ctx context.Context, req *order.SeckillPrePlaceOrderReq) (resp *order.SeckillPrePlaceOrderResp, err error) {
	resp, err = service.NewSeckillPrePlaceOrderService(ctx).Run(req)

	return resp, err
}

// QueryOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrder(ctx context.Context, req *order.QueryOrderReq) (resp *order.QueryOrderResp, err error) {
	resp, err = service.NewQueryOrderService(ctx).Run(req)

	return resp, err
}
