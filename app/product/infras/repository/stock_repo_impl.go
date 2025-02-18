package repository

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StockRepositoryImpl struct {
	db *gorm.DB
}

func (s *StockRepositoryImpl) DecrStock(ctx context.Context, productId, decr uint32) error {
	return s.updateStock(ctx, productId, decr, "decr")
}

func (s *StockRepositoryImpl) IncrStock(ctx context.Context, productId, incr uint32) error {
	return s.updateStock(ctx, productId, incr, "incr")
}

func (s *StockRepositoryImpl) updateStock(ctx context.Context, productId, change uint32, updateType string) error {
	productPOs := make([]*po.Product, 0)
	tx := s.db.Begin().WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}
	// todo add distributed lock 乐观锁
	if err := tx.Clauses(clause.Locking{Strength: "Update"}).Where("id = ?", productId).Find(&productPOs).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(productPOs) == 0 {
		tx.Rollback()
		return nil
	}
	productPO := productPOs[0]
	curStockNum := productPO.Stock
	if updateType == "incr" {
		curStockNum += change
	} else if updateType == "decr" {
		curStockNum -= change
	}
	if curStockNum < 0 {
		tx.Rollback()
		return nil
	}
	if err := tx.Model(&po.Product{}).Where("id = ?", productId).
		Updates(map[string]interface{}{
			"stock": curStockNum,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
