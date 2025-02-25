package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/cart/infra/rpc"
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteItemService struct {
	ctx context.Context
} // NewDeleteItemService new DeleteItemService
func NewDeleteItemService(ctx context.Context) *DeleteItemService {
	return &DeleteItemService{ctx: ctx}
}

// Run create note info
func (s *DeleteItemService) Run(req *cart.DeleteItemReq) (resp *cart.DeleteItemResp, err error) {
	// Finish your business logic.
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.ProductId})
	if err != nil {
		return nil, err
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	err = model.DeleteCartItem(mysql.DB, s.ctx, req.UserId, req.ProductId)

	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return
}
