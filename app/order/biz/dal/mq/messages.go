package mq

import "github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"

type PreOrderMessage struct {
	TempID    string `json:"temp_id"`
	UserID    uint32 `json:"user_id"`
	ProductID uint32 `json:"product_id"`
	Timestamp int64  `json:"timestamp"`
}

type DelayMessage struct {
	TempID     string `json:"temp_id"`
	Reason     string `json:"reason"` // timeout/stock_rollback
	CreatedAt  int64  `json:"created_at"`
	ExpectedAt int64  `json:"expected_at"`
}

type OrderMessage struct {
	TempID       string          `json:"temp_id"`
	OrderId      uint32          `json:"order_id"`
	UserCurrency string          `json:"user_currency"`
	Consignee    model.Consignee `json:"consignee"`
	ProductId    uint32          `json:"product_id"`
}
