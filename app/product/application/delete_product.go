package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	stateService := service.GetProductStateService()
	queryService := service.GetProductQueryService()
	updateService := service.GetProductUpdateService()

	origin, err := queryService.GetProductById(s.ctx, req.GetId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "queryService.GetProductById.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetProduct, "GetProductById failed")
	}
	validateFunc, err := stateService.GetCanTransferFunc(consts.ProductStatusDelete)
	if err != nil {
		klog.CtxErrorf(s.ctx, "GetCanTransferFunc.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrStateOperation, err.Error())
	}
	err = validateFunc(&service.ProductStateInfo{Status: origin.Status})
	if err != nil {
		klog.CtxErrorf(s.ctx, "validateFunc.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrStateOperation, err.Error())
	}

	target, err := stateService.ConstructTargetInfo(origin, consts.StateOperationTypeDel)
	if err != nil {
		klog.CtxErrorf(s.ctx, "ConstructTargetInfo.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrStateOperation, err.Error())
	}

	err = updateService.UpdateProduct(s.ctx, origin, target)
	if err != nil {
		klog.CtxErrorf(s.ctx, "updateService.UpdateProduct.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrUpdateProduct, "UpdateProduct failed")
	}

	return
}
