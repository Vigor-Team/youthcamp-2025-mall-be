package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type OfflineProductService struct {
	ctx context.Context
} // NewOfflineProductService new OfflineProductService
func NewOfflineProductService(ctx context.Context) *OfflineProductService {
	return &OfflineProductService{ctx: ctx}
}

// Run create note info
func (s *OfflineProductService) Run(req *product.OfflineProductReq) (resp *product.OfflineProductResp, err error) {
	stateService := service.GetProductStateService()
	queryService := service.GetProductQueryService()
	updateService := service.GetProductUpdateService()

	origin, err := queryService.GetProductById(context.Background(), req.GetId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "queryService.GetProductById.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetProduct, "GetProductById failed")
	}

	validateFunc, err := stateService.GetCanTransferFunc(consts.ProductStatusOffline)
	if err != nil {
		klog.CtxErrorf(s.ctx, "GetCanTransferFunc.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrStateOperation, err.Error())
	}
	err = validateFunc(&service.ProductStateInfo{Status: origin.Status})
	if err != nil {
		klog.CtxErrorf(s.ctx, "validateFunc.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrStateOperation, err.Error())
	}

	target, err := stateService.ConstructTargetInfo(origin, consts.StateOperationTypeOffline)
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
