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
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/application"

	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = application.NewListProductsService(ctx).Run(req)
	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = application.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = application.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// AddProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) AddProduct(ctx context.Context, req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	resp, err = application.NewAddProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = application.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = application.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// OnlineProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) OnlineProduct(ctx context.Context, req *product.OnlineProductReq) (resp *product.OnlineProductResp, err error) {
	resp, err = application.NewOnlineProductService(ctx).Run(req)

	return resp, err
}

// OfflineProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) OfflineProduct(ctx context.Context, req *product.OfflineProductReq) (resp *product.OfflineProductResp, err error) {
	resp, err = application.NewOfflineProductService(ctx).Run(req)

	return resp, err
}

// BatchGetProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) BatchGetProducts(ctx context.Context, req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	resp, err = application.NewBatchGetProductsService(ctx).Run(req)

	return resp, err
}

// GetCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetCategories(ctx context.Context, req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	resp, err = application.NewGetCategoriesService(ctx).Run(req)

	return resp, err
}

// DecrStock implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DecrStock(ctx context.Context, req *product.DecrStockReq) (resp *product.DecrStockResp, err error) {
	resp, err = application.NewDecrStockService(ctx).Run(req)

	return resp, err
}

// IncrStock implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) IncrStock(ctx context.Context, req *product.IncrStockReq) (resp *product.IncrStockResp, err error) {
	resp, err = application.NewIncrStockService(ctx).Run(req)

	return resp, err
}

// GetCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetCategory(ctx context.Context, req *product.GetCategoryReq) (resp *product.GetCategoryResp, err error) {
	resp, err = service.NewGetCategoryService(ctx).Run(req)

	return resp, err
}
