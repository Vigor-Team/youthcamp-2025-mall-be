package repository

import "context"

type StockRepository interface {
	DecrStock(ctx context.Context, productId, decr uint32) error
	IncrStock(ctx context.Context, productId, incr uint32) error
}
