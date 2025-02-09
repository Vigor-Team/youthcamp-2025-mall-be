package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/es"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository/converter"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (p *ProductRepositoryImpl) AddProduct(ctx context.Context, product *entity.ProductEntity) error {
	po, err := converter.ProductDO2POConverter.Convert2po(ctx, product)
	if err != nil {
		return err
	}
	go func() {
		pES := converter.ProductDo2ESConverter.Convert2Es(ctx, product)
		err := es.GetProductESClient().UpdateProduct(ctx, po.ID, pES)
		if err != nil {
			klog.CtxErrorf(ctx, "UpsertProductES err: %v", err)
		}
	}()
	return p.db.WithContext(ctx).Create(po).Error
}

func (p *ProductRepositoryImpl) UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	return nil
}

func (p *ProductRepositoryImpl) GetProductById(ctx context.Context, productId int64) (*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) ListProducts(ctx context.Context, filterParam map[string]interface{}) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	return nil, nil
}

func (p *ProductRepositoryImpl) SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	return nil, nil
}
