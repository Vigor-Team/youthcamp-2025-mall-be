package model

import "time"

type PreOrder struct {
	Base
	TempId    string `gorm:"uniqueIndex;size:256"`
	ProductId uint32
	UserId    uint32
	ExpiredAt time.Time
	Status    string
}

func (po PreOrder) TableName() string {
	return "pre_order"
}
