// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
