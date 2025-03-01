package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(_ *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(h.RequestContext.Value("user_id").(int32)),
	})
	if err != nil {
		return nil, err
	}

	items := make([]*cart.CartItem, 0, carts.Cart.Size())
	for _, i := range carts.Cart.Items {
		product, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: i.GetProductId(),
		})
		if err != nil {
			return nil, err
		}
		items = append(items, &cart.CartItem{
			ProductId:   i.GetProductId(),
			Name:        product.Product.Name,
			Price:       product.Product.Price,
			SpuName:     product.Product.SpuName,
			Description: product.Product.Description,
			Picture:     product.Product.Picture,
			SpuPrice:    product.Product.SpuPrice,
			Stock:       product.Product.Stock,
			Categories:  product.Product.Categories,
			Quantity:    i.GetQuantity(),
		})
	}

	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: carts.Cart.UserId,
			Items:  items,
		},
	}
	return
}
