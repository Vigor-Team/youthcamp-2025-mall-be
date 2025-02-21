package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"

	checkout "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCheckoutPreviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCheckoutPreviewService(Context context.Context, RequestContext *app.RequestContext) *GetCheckoutPreviewService {
	return &GetCheckoutPreviewService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCheckoutPreviewService) Run(_ *checkout.GetCheckoutPreviewReq) (resp *checkout.GetCheckoutPreviewResp, err error) {
	userId := gatewayutils.GetUserIdFromCtx(h.RequestContext)

	// Get cart items
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var cartNum int32
	var total float32
	items := make([]*checkout.Item, 0, carts.Cart.Size())
	for _, i := range carts.Cart.Items {
		// Get product info
		product, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: userId,
		})
		if err != nil {
			return nil, err
		}
		items = append(items, &checkout.Item{
			Name:     product.Product.Name,
			Price:    product.Product.Price,
			Picture:  product.Product.Picture,
			Quantity: i.GetQuantity(),
		})
		cartNum += 1
		total += float32(i.GetQuantity()) * product.Product.Price
	}
	resp = &checkout.GetCheckoutPreviewResp{
		CartNum: cartNum,
		Items:   items,
		Total:   total,
	}
	return
}
