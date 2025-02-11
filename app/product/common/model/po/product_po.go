package po

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint32     `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	SpuName     string     `json:"spu_name"`
	SpuPrice    float32    `json:"spu_price"`
	Price       float32    `json:"price"`
	Status      uint32     `json:"status"`
	Stock       uint32     `json:"stock"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p *Product) TableName() string {
	return "product"
}
