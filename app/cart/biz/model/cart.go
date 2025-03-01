package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (c Cart) TableName() string {
	return "cart"
}

func GetCartByUserId(db *gorm.DB, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}

func GetCartItemByUserIdAndProductId(db *gorm.DB, ctx context.Context, userId, productId uint32) (
	cart *Cart, err error,
) {
	var c Cart
	err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId, ProductId: productId}).First(&c).Error
	return &c, err
}

func UpdateCartQty(db *gorm.DB, ctx context.Context, userId, productId, qty uint32) error {
	return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId, ProductId: productId}).Update(
		"qty", qty,
	).Error
}

func DeleteCartItem(db *gorm.DB, ctx context.Context, userId, productId uint32) error {
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ? AND product_id = ?", userId, productId).Error
}

func AddCart(db *gorm.DB, ctx context.Context, c *Cart) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(
			&Cart{
				UserId: c.UserId, ProductId: c.ProductId,
			},
		).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}
	return err
}

func EmptyCart(db *gorm.DB, ctx context.Context, userId uint32) error {
	if userId == 0 {
		return errors.New("user_is is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}
