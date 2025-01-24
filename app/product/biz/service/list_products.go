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

package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/biz/model"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	c, err := model.ListProducts(mysql.DB, s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		resp.Products = append(resp.Products, &product.Product{Id: uint32(v1.ID), Name: v1.Name, Description: v1.Description, Picture: v1.Picture, Price: v1.Price, Quantity: v1.Quantity})
	}

	return resp, nil
}
