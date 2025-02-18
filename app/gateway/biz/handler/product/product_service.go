package product

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateProduct .
// @router /api/v1/products [POST]
func CreateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewCreateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// UpdateProduct .
// @router /api/v1/products/:productId [PUT]
func UpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.UpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// DeleteProduct .
// @router /api/v1/products/:productId [DELETE]
func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewDeleteProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// OnlineProduct .
// @router /api/v1/products/:productId/online [POST]
func OnlineProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewOnlineProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// OfflineProduct .
// @router /api/v1/products/:productId/offline [POST]
func OfflineProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewOfflineProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// DecreaseStock .
// @router /api/v1/products/:productId/stock/decrease [PATCH]
func DecreaseStock(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.StockOpReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewDecreaseStockService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// IncreaseStock .
// @router /api/v1/products/:productId/stock/increase [PATCH]
func IncreaseStock(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.StockOpReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewIncreaseStockService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// GetProduct .
// @router /api/v1/products/:productId [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.GetProductResp{}
	resp, err = service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// ListProducts .
// @router /api/v1/products [GET]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.ListProductsResp{}
	resp, err = service.NewListProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// BatchGetProducts .
// @router /api/v1/products/batch [GET]
func BatchGetProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.BatchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.BatchProductsResp{}
	resp, err = service.NewBatchGetProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// SearchProducts .
// @router /api/v1/search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.ListProductsResp{}
	resp, err = service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// ListCategories .
// @router /api/v1/categories [GET]
func ListCategories(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.ListCategoriesResp{}
	resp, err = service.NewListCategoriesService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}

// GetCategory .
// @router /api/v1/categories/:categoryId [GET]
func GetCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetCategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	resp := &product.GetCategoryResp{}
	resp, err = service.NewGetCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.FailResponse(ctx, c, err)
		return
	}

	utils.SuccessResponse(c, resp)
}
