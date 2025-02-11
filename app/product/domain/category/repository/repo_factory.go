package repository

type CategoryRepositoryFactory struct {
	categoryRepository CategoryRepository
}

var inst = &CategoryRepositoryFactory{}

func GetFactory() *CategoryRepositoryFactory {
	return inst
}

func (f *CategoryRepositoryFactory) GetCategoryRepository() CategoryRepository {
	return f.categoryRepository
}

func (f *CategoryRepositoryFactory) SetCategoryRepository(categoryRepositoryIns CategoryRepository) {
	f.categoryRepository = categoryRepositoryIns
}
