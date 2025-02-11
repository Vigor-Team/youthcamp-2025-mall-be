package po

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

func (c *Category) TableName() string {
	return "category"
}
