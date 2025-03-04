package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

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
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "quantity must be greater than 0")
	}

	// Check if the product exists
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		klog.CtxErrorf(s.ctx, "rpc.ProductClient.GetProduct.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "rpc.ProductClient.GetProduct error")
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "product not exist")
	}

	// Check if the cart exists and has the product
	cartItem, err := model.GetCartItemByUserIdAndProductId(mysql.DB, s.ctx, req.UserId, req.Item.ProductId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.GetCartItemByUserIdAndProductId.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUpdateCart, "get cart item error")
	}
	if cartItem == nil {
		return nil, kerrors.NewBizStatusError(consts.ErrUpdateCart, "item not found in cart")
	}

	// Update the quantity of the product in the cart
	err = model.UpdateCartQty(
		mysql.DB, s.ctx, req.UserId, req.Item.ProductId, uint32(req.Item.Quantity),
	)

	if err != nil {
		klog.CtxErrorf(s.ctx, "model.UpdateCartQty.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUpdateCart, "update cart item error")
	}

	return &cart.UpdateCartResp{}, nil
}
