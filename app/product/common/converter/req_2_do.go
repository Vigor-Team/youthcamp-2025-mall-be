package converter

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/constant"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
)

func ConvertAddReq2Entity(_ context.Context, req *product.AddProductReq) (*entity.ProductEntity, error) {
	pid, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}
	return &entity.ProductEntity{
		ID:          uint32(pid),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SpuName:     req.SpuName,
		SpuPrice:    req.SpuPrice,
		Picture:     req.Picture,
		Status:      constant.ProductStatusOnline,
	}, nil
}

func ConvertUpdateReq2Entity(_ context.Context, req *product.UpdateProductReq) (*entity.ProductEntity, error) {
	return &entity.ProductEntity{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SpuName:     req.SpuName,
		SpuPrice:    req.SpuPrice,
		Picture:     req.Picture,
		Status:      constant.ProductStatusOnline,
	}, nil
}
