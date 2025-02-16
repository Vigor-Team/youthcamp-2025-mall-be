package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PreOrder struct {
	Base
	ProductId uint32
	UserId    uint32
	Status    string
	ExpiredAt time.Time
}

func (po PreOrder) TableName() string {
	return "pre_order"
}

func AddPreOrder(db *gorm.DB, ctx context.Context, po *PreOrder) error {
	return db.Model(&PreOrder{}).Create(po).Error
}
