package strategy

import "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/constant"

type ListingStrategy interface {
	AllowedStatuses() []consts.ProductStatus
	Filter() map[string]interface{}
}

type CustomerStrategy struct {
	FilterParam map[string]interface{}
}

func (s *CustomerStrategy) AllowedStatuses() []consts.ProductStatus {
	return []consts.ProductStatus{consts.ProductStatusOnline}
}

func (s *CustomerStrategy) Filter() map[string]interface{} {
	return s.FilterParam
}

type MerchantStrategy struct {
	FilterParam map[string]interface{}
}

func (s *MerchantStrategy) Filter() map[string]interface{} {
	return s.FilterParam
}

func (s *MerchantStrategy) AllowedStatuses() []consts.ProductStatus {
	return []consts.ProductStatus{consts.ProductStatusOffline, consts.ProductStatusOnline}
}

type AdminStrategy struct {
	FilterParam map[string]interface{}
}

func (s *AdminStrategy) Filter() map[string]interface{} {
	return s.FilterParam
}

func (s *AdminStrategy) AllowedStatuses() []consts.ProductStatus {
	return []consts.ProductStatus{consts.ProductStatusOnline, consts.ProductStatusDelete, consts.ProductStatusOffline}
}

func NewListingStrategy(role string, filterParam map[string]interface{}) ListingStrategy {
	switch role {
	case "Seller":
		return &MerchantStrategy{FilterParam: filterParam}
	case "Admin":
		return &AdminStrategy{FilterParam: filterParam}
	case "Super Admin":
		return &AdminStrategy{FilterParam: filterParam}
	default:
		return &CustomerStrategy{FilterParam: filterParam}
	}
}
