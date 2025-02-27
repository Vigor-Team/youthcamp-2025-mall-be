package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type UpdateCartService struct {
	ctx context.Context
} // NewUpdateCartService new UpdateCartService
func NewUpdateCartService(ctx context.Context) *UpdateCartService {
	return &UpdateCartService{ctx: ctx}
}

// Run create note info
func (s *UpdateCartService) Run(req *cart.UpdateCartReq) (resp *cart.UpdateCartResp, err error) {
	// Finish your business logic.
	if req.Item.Quantity < 0 {
		return nil, kerrors.NewBizStatusError(40003, "quantity must be greater than 0")
	}

	// Check if the product exists
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		return nil, err
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	// Check if the cart exists and has the product
	cartItem, err := model.GetCartItemByUserIdAndProductId(mysql.DB, s.ctx, req.UserId, req.Item.ProductId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40010, err.Error())
	}
	if cartItem == nil {
		return nil, kerrors.NewBizStatusError(40005, "item not found in cart")
	}

	// Update the quantity of the product in the cart
	err = model.UpdateCartQty(
		mysql.DB, s.ctx, req.UserId, req.Item.ProductId, uint32(req.Item.Quantity),
	)

	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.UpdateCartResp{}, nil
}
