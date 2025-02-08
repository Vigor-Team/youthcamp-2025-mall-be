package repository

type ProductRepositoryFactory struct {
	productRepository ProductRepository
	stockRepository   StockRepository
}

var inst = &ProductRepositoryFactory{}

func GetFactory() *ProductRepositoryFactory {
	return inst
}

func (f *ProductRepositoryFactory) GetProductRepository() ProductRepository {
	return f.productRepository
}

func (f *ProductRepositoryFactory) SetProductRepository(productRepositoryIns ProductRepository) {
	f.productRepository = productRepositoryIns
}

func (f *ProductRepositoryFactory) GetStockRepository() StockRepository {
	return f.stockRepository
}

func (f *ProductRepositoryFactory) SetStockRepository(stockRepositoryIns StockRepository) {
	f.stockRepository = stockRepositoryIns
}
