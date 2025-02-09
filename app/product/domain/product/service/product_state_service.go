package service

import (
	"context"
	"errors"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/constant"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	productrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/repository"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ProductStateService struct {
}

var productStateService ProductStateService

func GetProductStateService() *ProductStateService {
	return &productStateService
}

type ProductStateInfo struct {
	Status constant.ProductStatus
}

type CanTransferFunc func(originalInfo *ProductStateInfo) error

type ConstructTargetInfoFunc func(originalInfo *ProductStateInfo) *ProductStateInfo

var canTransferFuncMap = map[constant.StateOperationType]CanTransferFunc{
	constant.StateOperationTypeAdd: func(originalInfo *ProductStateInfo) error {
		return nil
	},
	constant.StateOperationTypeSave: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status == constant.ProductStatusDelete {
			return errors.New("product has been deleted")
		}
		return nil
	},
	constant.StateOperationTypeDel: func(originalInfo *ProductStateInfo) error {
		return nil
	},
	constant.StateOperationTypeOnline: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status != constant.ProductStatusOffline {
			return errors.New("product is not offline")
		}
		return nil
	},
	constant.StateOperationTypeOffline: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status != constant.ProductStatusOnline {
			return errors.New("product is not online")
		}
		return nil
	},
}

var constructTargetInfoFuncMap = map[constant.StateOperationType]ConstructTargetInfoFunc{
	constant.StateOperationTypeAdd: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeSave: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeDel: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusDelete
		return
	},
	constant.StateOperationTypeOnline: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeOffline: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOffline
		return
	},
}

func (s *ProductStateService) GetCanTransferFunc(operationType constant.StateOperationType) (CanTransferFunc, error) {
	if canTransferFunc, ok := canTransferFuncMap[operationType]; ok {
		return canTransferFunc, nil
	}
	return nil, errors.New("GetCanTransferFunc not found")
}

func (s *ProductStateService) ConstructTargetInfo(originProduct *entity.ProductEntity,
	operation constant.StateOperationType,
) (*entity.ProductEntity, error) {
	targetProduct, err := originProduct.Clone()
	if err != nil {
		return nil, err
	}
	originStateInfo := &ProductStateInfo{
		Status: originProduct.Status,
	}
	constructFunc, err := s.getConstructTargetInfoFunc(operation)
	if err != nil {
		return nil, err
	}
	targetState := constructFunc(originStateInfo)
	targetProduct.Status = targetState.Status
	return targetProduct, nil
}

func (s *ProductStateService) OperateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	err := productrepo.GetFactory().GetProductRepository().UpdateProduct(ctx, origin, target)
	if err != nil {
		klog.CtxErrorf(ctx, "OperateProduct err: %v", err)
		return err
	}
	return nil
}

func (s *ProductStateService) getConstructTargetInfoFunc(operationType constant.StateOperationType) (ConstructTargetInfoFunc, error) {
	if constructTargetInfoFunc, ok := constructTargetInfoFuncMap[operationType]; ok {
		return constructTargetInfoFunc, nil
	}
	return nil, errors.New("GetConstructTargetInfoFunc not found")
}
