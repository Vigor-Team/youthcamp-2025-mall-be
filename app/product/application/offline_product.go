package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/constant"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
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
		return nil, err
	}

	validateFunc, err := stateService.GetCanTransferFunc(constant.ProductStatusOffline)
	if err != nil {
		return nil, err
	}
	err = validateFunc(&service.ProductStateInfo{Status: origin.Status})
	if err != nil {
		return nil, err
	}

	target, err := stateService.ConstructTargetInfo(origin, constant.StateOperationTypeOffline)
	if err != nil {
		return nil, err
	}

	err = updateService.UpdateProduct(s.ctx, origin, target)
	if err != nil {
		return nil, err
	}

	return
}
