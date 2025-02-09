package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/constant"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/service"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

type OnlineProductService struct {
	ctx context.Context
} // NewOnlineProductService new OnlineProductService
func NewOnlineProductService(ctx context.Context) *OnlineProductService {
	return &OnlineProductService{ctx: ctx}
}

// Run create note info
func (s *OnlineProductService) Run(req *product.OnlineProductReq) (resp *product.OnlineProductResp, err error) {
	stateService := service.GetProductStateService()
	queryService := service.GetProductQueryService()
	updateService := service.GetProductUpdateService()

	origin, err := queryService.GetProductById(context.Background(), req.GetId())
	if err != nil {
		return nil, err
	}

	validateFunc, err := stateService.GetCanTransferFunc(constant.ProductStatusOnline)
	if err != nil {
		return nil, err
	}
	err = validateFunc(&service.ProductStateInfo{Status: origin.Status})
	if err != nil {
		return nil, err
	}

	target, err := stateService.ConstructTargetInfo(origin, constant.StateOperationTypeOnline)
	if err != nil {
		return nil, err
	}

	err = updateService.UpdateProduct(s.ctx, origin, target)
	if err != nil {
		return nil, err
	}

	return
}
