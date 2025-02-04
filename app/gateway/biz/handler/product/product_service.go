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

package product

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetProduct .
// @router /products/:productId [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	hlog.CtxInfof(ctx, "GetProduct: %v", resp)
	utils.SuccessResponse(c, resp)
}

// SearchProducts .
// @router /search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	utils.SuccessResponse(c, resp)
}

// ListProducts .
// @router /products [GET]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewListProductsService(ctx, c).Run(&req)

	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}

// ListCategories .
// @router /categories [GET]
func ListCategories(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewListCategoriesService(ctx, c).Run(&req)

	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	utils.SuccessResponse(c, resp)
}

// GetCategory .
// @router /categories/:categoryId [GET]
func GetCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp, err := service.NewGetCategoryService(ctx, c).Run(&req)

	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	utils.SuccessResponse(c, resp)
}
