package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/strategy"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/embedding"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/es"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository/converter"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository/differ"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (p *ProductRepositoryImpl) AddProduct(ctx context.Context, product *entity.ProductEntity) error {
	productPO, err := converter.ProductDO2POConverter.Convert2po(ctx, product)
	if err != nil {
		return err
	}
	eb, err := embedding.GetArkEmbedding(ctx, nil)
	if err != nil {
		return err
	}
	go func() {
		pES := converter.ProductDoWithESConverter.Convert2ES(ctx, product)
		texts := []string{
			pES.Name,
			pES.Description,
			pES.SpuName,
		}
		if pES.CategoryNames != nil {
			texts = append(texts, pES.CategoryNames...)
		}
		vectors, err := eb.EmbedStrings(ctx, texts)
		if err != nil {
			klog.CtxErrorf(ctx, "EmbedStrings err: %v", err)
			return
		}
		var merged []float32
		if len(vectors) > 0 {
			merged = utils.MergeVectors(vectors)
		}
		pES.Embedding = merged
		err = es.GetProductESClient().UpsertProduct(ctx, productPO.ID, pES)
		if err != nil {
			klog.CtxErrorf(ctx, "UpsertProductES.err: %v", err)
		}
	}()
	return p.db.WithContext(ctx).Create(productPO).Error
}

func (p *ProductRepositoryImpl) UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	productId := target.ID
	originPO, err := converter.ProductDO2POConverter.Convert2po(ctx, origin)
	if err != nil {
		return err
	}
	targetPO, err := converter.ProductDO2POConverter.Convert2po(ctx, target)
	if err != nil {
		return err
	}
	eb, err := embedding.GetArkEmbedding(ctx, nil)
	if err != nil {
		return err
	}
	// todo 数据一致性
	go func() {
		pES := converter.ProductDoWithESConverter.Convert2ES(ctx, target)
		texts := []string{
			pES.Name,
			pES.Description,
			pES.SpuName,
		}
		if pES.CategoryNames != nil {
			texts = append(texts, pES.CategoryNames...)
		}
		vectors, err := eb.EmbedStrings(ctx, texts)
		if err != nil {
			klog.CtxErrorf(ctx, "EmbedStrings err: %v", err)
			return
		}
		var merged []float32
		if len(vectors) > 0 {
			merged = utils.MergeVectors(vectors)
		}
		pES.Embedding = merged
		err = es.GetProductESClient().UpsertProduct(ctx, productId, pES)
		if err != nil {
			klog.CtxErrorf(ctx, "UpsertProductES err: %v", err)
		}
	}()
	changeMap := differ.ProductPODiffer.GetChangedMap(originPO, targetPO)
	fmt.Println("changeMap: ", changeMap)
	return DB.WithContext(ctx).Model(&po.Product{}).Where("id = ?", productId).
		Updates(changeMap).Error
}

func (p *ProductRepositoryImpl) GetProductById(ctx context.Context, productId uint32) (*entity.ProductEntity, error) {
	products := make([]*po.Product, 0)
	err := p.db.WithContext(ctx).Preload("Categories").Where("id = ?", productId).First(&products).Error
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("product not found")
	}
	do, err := converter.ProductPO2DOConverter.Convert2do(ctx, products[0])
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (p *ProductRepositoryImpl) ListProducts(ctx context.Context, strategy strategy.ListingStrategy) ([]*entity.ProductEntity, error) {
	filterParam := strategy.Filter()
	products := make([]*po.Product, 0)
	productEntities := make([]*entity.ProductEntity, 0)

	db := p.db.Debug().WithContext(ctx)

	if categoryId, exists := filterParam["category_id"]; exists {
		delete(filterParam, "category_id")
		db = db.Joins("JOIN product_category ON product_category.product_id = product.id").
			Joins("JOIN category ON category.id = product_category.category_id").
			Where("category.id = ?", categoryId)
	}

	allowedStatuses := strategy.AllowedStatuses()
	if len(allowedStatuses) > 0 {
		db = db.Where("product.status IN ?", allowedStatuses)
	}

	for k, v := range filterParam {
		db = db.Where(k+" = ?", v)
	}
	if err = db.Scopes(AvailableProducts).Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}
	for _, product := range products {
		do, err := converter.ProductPO2DOConverter.Convert2do(ctx, product)
		if err != nil {
			return nil, err
		}
		productEntities = append(productEntities, do)
	}
	return productEntities, nil
}

func (p *ProductRepositoryImpl) BatchGetProducts(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	entities, err := es.GetProductESClient().BatchGetProductById(ctx, productIds)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (p *ProductRepositoryImpl) SearchProducts(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	entities, err := es.GetProductESClient().SearchProduct(ctx, keyword)
	if err != nil {
		return nil, err
	}
	return entities, nil
}
