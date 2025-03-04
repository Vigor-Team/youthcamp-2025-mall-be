package mq

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/redis"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func handlePreOrder(ctx context.Context, msg PreOrderMessage) error {
	// get distributed lock
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 20*time.Second); !success {
		klog.CtxErrorf(ctx, "redis.TryLock.err")
		return fmt.Errorf("redis.TryLock.err")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "redis.ReleaseLock.err: %v", err)
		}
	}(ctx, key)

	// add preorder
	id, err := strconv.Atoi(msg.TempID)
	if err != nil {
		return err
	}
	if err = model.AddPreOrder(mysql.DB, ctx, &model.PreOrder{
		Base: model.Base{
			ID:        uint32(id),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ProductId: msg.ProductID,
		UserId:    msg.UserID,
		Status:    "pending",
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}); err != nil {
		return err
	}

	// publish delay message
	producer := NewProducer(Client)
	delayMsg := DelayMessage{
		TempID:     msg.TempID,
		UserID:     msg.UserID,
		ProductID:  msg.ProductID,
		CreatedAt:  time.Now().Unix(),
		ExpectedAt: time.Now().Add(10 * time.Minute).Unix(),
	}
	if err = producer.PublishDelay(ctx, delayMsg, 1*time.Second); err != nil {
		return err
	}
	return nil
}

func handleOrder(ctx context.Context, msg OrderMessage) error {
	// get distributed lock
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 20*time.Second); !success {
		klog.CtxErrorf(ctx, "redis.TryLock.err")
		return fmt.Errorf("redis.TryLock.err")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "redis.ReleaseLock.err: %v", err)
		}
	}(ctx, key)

	// check order exists
	_, err := model.GetOrder(mysql.DB, ctx, msg.UserID, strconv.Itoa(int(msg.OrderId)))
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.CtxErrorf(ctx, "model.GetOrder.err: %v", err)
		return fmt.Errorf("model.GetOrder.err: %v", err)
	}

	// start transaction
	tx := mysql.DB.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	defer tx.Rollback()

	// update pre order
	if err = tx.Model(&model.PreOrder{}).Where("id = ?", msg.TempID).Update("status", "completed").Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Model.Update.err: %v", err)
		return err
	}

	// decrease product stock
	_, err = rpc.ProductClient.DecrStock(ctx, &rpcproduct.DecrStockReq{
		Id:   msg.ProductId,
		Decr: 1,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "rpc.ProductClient.DecrStock.err: %v", err)
		return err
	}

	// create order
	o := &model.Order{
		OrderId:      strconv.Itoa(int(msg.OrderId)),
		OrderState:   model.OrderStatePlaced,
		UserId:       msg.UserID,
		UserCurrency: msg.UserCurrency,
		Consignee: model.Consignee{
			Email:         msg.Consignee.Email,
			StreetAddress: msg.Consignee.StreetAddress,
			City:          msg.Consignee.City,
			State:         msg.Consignee.State,
			Country:       msg.Consignee.Country,
			ZipCode:       msg.Consignee.ZipCode,
		},
	}
	if err := tx.Create(o).Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Create.err: %v", err)
		return err
	}

	// create order item
	var itemList []*model.OrderItem
	itemList = append(itemList, &model.OrderItem{
		OrderIdRefer: o.OrderId,
		ProductId:    msg.ProductId,
		Quantity:     1,
		Cost:         msg.Cost,
	})
	if err = tx.Create(&itemList).Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Create.err: %v", err)
		return err
	}

	// delete redis key
	preOrderID, err := strconv.ParseUint(msg.TempID, 10, 32)
	if err != nil {
		klog.CtxErrorf(ctx, "strconv.ParseUint.err: %v", err)
		return err
	}
	preOrderKey := redis.GetOrderPreOrderKey(uint32(preOrderID))
	if err = redis.RedisClient.HDel(ctx, preOrderKey, "user_id", "product_id").Err(); err != nil {
		klog.CtxErrorf(ctx, "redis.RedisClient.HDel.err: %v", err)
		return err
	}

	// commit transaction
	if err = tx.Commit().Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Commit.err: %v", err)
		return err
	}
	return nil
}

func handleDelayOrder(ctx context.Context, msg DelayMessage) error {
	// get distributed lock
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 10); !success {
		klog.CtxErrorf(ctx, "redis.TryLock.err")
		return fmt.Errorf("redis.TryLock.err")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "redis.ReleaseLock.err: %v", err)
		}
	}(ctx, key)

	// check order
	var order model.Order
	if err := mysql.DB.Where("pre_order_id = ?", msg.TempID).First(&order).Error; err != nil {
		if order.OrderState == model.OrderStatePlaced {
			return nil
		}
	}

	tx := mysql.DB.Begin()
	defer tx.Rollback()

	// get pre order
	var preOrder model.PreOrder
	if err := mysql.DB.Where("id = ?", msg.TempID).First(&preOrder).Error; err != nil {
		klog.CtxErrorf(ctx, "mysql.DB.Where.err: %v", err)
		return err
	}

	// update pre order status
	if err := tx.Model(&model.PreOrder{}).Where("id = ?", msg.TempID).Update("status", "cancelled").Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Model.Update.err: %v", err)
		return err
	}

	// update order status
	if err := tx.Model(&model.Order{}).Where("pre_order_id = ?", msg.TempID).Update("order_state", model.OrderStateCanceled).Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Model.Update.err: %v", err)
		return err
	}

	// recover product stock
	_, err := rpc.ProductClient.IncrStock(ctx, &rpcproduct.IncrStockReq{
		Id:   preOrder.ProductId,
		Incr: 1,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "rpc.ProductClient.IncrStock.err: %v", err)
		return err
	}

	// rollback redis
	preOrderID, err := strconv.ParseUint(msg.TempID, 10, 32)
	if err != nil {
		klog.CtxErrorf(ctx, "strconv.ParseUint.err: %v", err)
		return err
	}
	preOrderKey := redis.GetOrderPreOrderKey(uint32(preOrderID))
	productOrderKey := redis.GetProductOrderKey(preOrder.ProductId)
	stockKey := redis.GetProductStockKey(msg.ProductID)

	if err = redis.RedisClient.SRem(ctx, productOrderKey, msg.UserID).Err(); err != nil {
		klog.CtxErrorf(ctx, "redis.RedisClient.SRem.err: %v", err)
		return err
	}
	if err = redis.RedisClient.HDel(ctx, preOrderKey, "user_id", "product_id").Err(); err != nil {
		klog.CtxErrorf(ctx, "redis.RedisClient.HDel.err: %v", err)
		return err
	}

	if err = redis.RedisClient.Incr(ctx, stockKey).Err(); err != nil {
		klog.CtxErrorf(ctx, "redis.RedisClient.Incr.err: %v", err)
		return err
	}

	// commit transaction
	if err = tx.Commit().Error; err != nil {
		klog.CtxErrorf(ctx, "tx.Commit.err: %v", err)
		return err
	}

	return nil
}
