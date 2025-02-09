package repository

import "gorm.io/gorm"

// todo opt scope
func AvailableProducts(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", 0)
}
