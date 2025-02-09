package product

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddProduct(ctx context.Context, req *product.AddProductReq, callOptions ...callopt.Option) (resp *product.AddProductResp, err error) {
	resp, err = defaultClient.AddProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.UpdateProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func OnlineProduct(ctx context.Context, req *product.OnlineProductReq, callOptions ...callopt.Option) (resp *product.OnlineProductResp, err error) {
	resp, err = defaultClient.OnlineProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "OnlineProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func OfflineProduct(ctx context.Context, req *product.OfflineProductReq, callOptions ...callopt.Option) (resp *product.OfflineProductResp, err error) {
	resp, err = defaultClient.OfflineProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "OfflineProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BatchGetProducts(ctx context.Context, req *product.BatchGetProductsReq, callOptions ...callopt.Option) (resp *product.BatchGetProductsResp, err error) {
	resp, err = defaultClient.BatchGetProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchGetProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategories(ctx context.Context, req *product.GetCategoriesReq, callOptions ...callopt.Option) (resp *product.GetCategoriesResp, err error) {
	resp, err = defaultClient.GetCategories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DecrStock(ctx context.Context, req *product.DecrStockReq, callOptions ...callopt.Option) (resp *product.DecrStockResp, err error) {
	resp, err = defaultClient.DecrStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DecrStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func IncrStock(ctx context.Context, req *product.IncrStockReq, callOptions ...callopt.Option) (resp *product.IncrStockResp, err error) {
	resp, err = defaultClient.IncrStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "IncrStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
