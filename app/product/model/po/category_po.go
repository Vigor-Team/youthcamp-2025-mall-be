package po

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}
