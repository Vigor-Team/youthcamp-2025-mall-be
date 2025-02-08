package product

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	OnlineProduct(ctx context.Context, Req *product.OnlineProductReq, callOptions ...callopt.Option) (r *product.OnlineProductResp, err error)
	OfflineProduct(ctx context.Context, Req *product.OfflineProductReq, callOptions ...callopt.Option) (r *product.OfflineProductResp, err error)
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	BatchGetProducts(ctx context.Context, Req *product.BatchGetProductsReq, callOptions ...callopt.Option) (r *product.BatchGetProductsResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	GetCategories(ctx context.Context, Req *product.GetCategoriesReq, callOptions ...callopt.Option) (r *product.GetCategoriesResp, err error)
	DecrStock(ctx context.Context, Req *product.DecrStockReq, callOptions ...callopt.Option) (r *product.DecrStockResp, err error)
	IncrStock(ctx context.Context, Req *product.IncrStockReq, callOptions ...callopt.Option) (r *product.IncrStockResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error) {
	return c.kitexClient.AddProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) OnlineProduct(ctx context.Context, Req *product.OnlineProductReq, callOptions ...callopt.Option) (r *product.OnlineProductResp, err error) {
	return c.kitexClient.OnlineProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) OfflineProduct(ctx context.Context, Req *product.OfflineProductReq, callOptions ...callopt.Option) (r *product.OfflineProductResp, err error) {
	return c.kitexClient.OfflineProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) BatchGetProducts(ctx context.Context, Req *product.BatchGetProductsReq, callOptions ...callopt.Option) (r *product.BatchGetProductsResp, err error) {
	return c.kitexClient.BatchGetProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCategories(ctx context.Context, Req *product.GetCategoriesReq, callOptions ...callopt.Option) (r *product.GetCategoriesResp, err error) {
	return c.kitexClient.GetCategories(ctx, Req, callOptions...)
}

func (c *clientImpl) DecrStock(ctx context.Context, Req *product.DecrStockReq, callOptions ...callopt.Option) (r *product.DecrStockResp, err error) {
	return c.kitexClient.DecrStock(ctx, Req, callOptions...)
}

func (c *clientImpl) IncrStock(ctx context.Context, Req *product.IncrStockReq, callOptions ...callopt.Option) (r *product.IncrStockResp, err error) {
	return c.kitexClient.IncrStock(ctx, Req, callOptions...)
}
