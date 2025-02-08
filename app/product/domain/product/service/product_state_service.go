package service

type ProductStateService struct {
}

var productStateService ProductStateService

func GetProductStateService() *ProductStateService {
	return &productStateService
}
