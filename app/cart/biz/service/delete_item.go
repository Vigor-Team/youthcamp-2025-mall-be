package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteItemService struct {
	ctx context.Context
} // NewDeleteItemService new DeleteItemService
func NewDeleteItemService(ctx context.Context) *DeleteItemService {
	return &DeleteItemService{ctx: ctx}
}

// Run create note info
func (s *DeleteItemService) Run(req *cart.DeleteItemReq) (resp *cart.DeleteItemResp, err error) {
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.ProductId})
	if err != nil {
		klog.CtxErrorf(s.ctx, "rpc.ProductClient.GetProduct.err: %v", err)
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "rpc.ProductClient.GetProduct error")
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "product not exist")
	}

	err = model.DeleteCartItem(mysql.DB, s.ctx, req.UserId, req.ProductId)

	if err != nil {
		klog.CtxErrorf(s.ctx, "model.DeleteCartItem.err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrDeleteCart, "delete cart item error")
	}

	return
}
