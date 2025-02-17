package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PreOrder struct {
	Base
	ProductId uint32    `json:"product_id"`
	UserId    uint32    `json:"user_id"`
	Status    string    `json:"status"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (po PreOrder) TableName() string {
	return "pre_order"
}

func AddPreOrder(db *gorm.DB, ctx context.Context, po *PreOrder) error {
	return db.Model(&PreOrder{}).Create(po).Error
}
