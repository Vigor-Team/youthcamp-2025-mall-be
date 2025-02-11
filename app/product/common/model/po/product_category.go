package po

type ProductCategory struct {
	ProductID  uint32 `json:"product_id"`
	CategoryID uint32 `json:"category_id"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}
