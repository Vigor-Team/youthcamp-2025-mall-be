package service

import (
	"context"

	common "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/common"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.AddProduct(h.Context, &rpcproduct.AddProductReq{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Picture:     req.Picture,
		SpuPrice:    req.SpuPrice,
		SpuName:     req.SpuName,
		Stock:       req.Stock,
		CategoryIds: req.CategoryIds,
	})
	if err != nil {
		return
	}
	return
}
