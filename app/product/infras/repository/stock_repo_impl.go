package repository

import (
	"context"
	"gorm.io/gorm"
)

type StockRepositoryImpl struct {
	db *gorm.DB
}

func (s *StockRepositoryImpl) DecrStock(ctx context.Context, productId, decr uint32) error {
	return nil
}

func (s *StockRepositoryImpl) IncrStock(ctx context.Context, productId, incr uint32) error {
	return nil
}
