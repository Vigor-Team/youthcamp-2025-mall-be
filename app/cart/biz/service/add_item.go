package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	if req.Item.Quantity < 0 {
		klog.CtxErrorf(s.ctx, "quantity must be greater than 0")
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "quantity must be greater than 0")
	}

	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		klog.CtxErrorf(s.ctx, "rpc.ProductClient.GetProduct.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "rpc.ProductClient.GetProduct error")
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "product not exist")
	}

	err = model.AddCart(
		mysql.DB, s.ctx, &model.Cart{
			UserId:    req.UserId,
			ProductId: req.Item.ProductId,
			Qty:       uint32(req.Item.Quantity),
		},
	)
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.AddCart.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrAddCart, "add cart error")
	}

	return &cart.AddItemResp{}, nil
}
