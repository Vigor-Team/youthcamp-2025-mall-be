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
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 10); !success {
		return errors.New("get lock failed")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "release lock failed: %v", err)
		}
	}(ctx, key)

	id, err := strconv.Atoi(msg.TempID)
	if err != nil {
		return err
	}
	if err = model.AddPreOrder(mysql.DB, ctx, &model.PreOrder{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ProductId: msg.ProductID,
		UserId:    msg.UserID,
		Status:    "pending",
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}); err != nil {
		return errors.New("add pre order failed")
	}

	producer := NewProducer(Client)
	delayMsg := DelayMessage{
		TempID:     msg.TempID,
		CreatedAt:  time.Now().Unix(),
		ExpectedAt: time.Now().Add(10 * time.Minute).Unix(),
	}
	if err = producer.PublishDelay(ctx, delayMsg, 10*time.Minute); err != nil {
		return errors.New("publish delay message failed")
	}
	return nil
}

func handleOrder(ctx context.Context, msg OrderMessage) error {
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 10); !success {
		return errors.New("get lock failed")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "release lock failed: %v", err)
		}
	}(ctx, key)

	_, err := model.GetOrder(mysql.DB, ctx, msg.UserID, strconv.Itoa(int(msg.OrderId)))
	if err != nil {
		return err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("order already exists")
	}

	tx := mysql.DB.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	defer tx.Rollback()

	// update pre order
	if err := tx.Model(&model.PreOrder{}).Where("id = ?", msg.OrderId).Update("status", "completed").Error; err != nil {
		return fmt.Errorf("update pre order failed: %w", err)
	}

	// decrease product stock
	_, err = rpc.ProductClient.DecrStock(ctx, &rpcproduct.DecrStockReq{
		Id:   msg.ProductId,
		Decr: 1,
	})
	if err != nil {
		return fmt.Errorf("decrease product stock failed: %w", err)
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
	var itemList []*model.OrderItem
	itemList = append(itemList, &model.OrderItem{
		OrderIdRefer: o.OrderId,
		ProductId:    msg.ProductId,
		Quantity:     1,
		Cost:         msg.Cost,
	})
	if err := tx.Create(&itemList).Error; err != nil {
		return err
	}

	// delete redis temp key
	tempKey := redis.GetSeckillTempKey(msg.TempID)
	err = redis.RedisClient.Del(ctx, tempKey).Err()
	if err != nil {
		return fmt.Errorf("delete temp key failed: %w", err)
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %w", err)
	}
	return nil
}

func handleDelayOrder(ctx context.Context, msg DelayMessage) error {
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 10); !success {
		return errors.New("get lock failed")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "release lock failed: %v", err)
		}
	}(ctx, key)

	// 查询订单
	var order model.Order
	if err := mysql.DB.Where("pre_order_id = ?", msg.TempID).First(&order).Error; err != nil {
		if order.OrderState == model.OrderStatePlaced {
			return nil
		}
	}

	tx := mysql.DB.Begin()
	defer tx.Rollback()

	// 查询pre_order表
	var preOrder model.PreOrder
	if err := mysql.DB.Where("id = ?", msg.TempID).First(&preOrder); err != nil {
		return fmt.Errorf("get pre order failed: %w", err)
	}

	// 恢复库存
	_, err := rpc.ProductClient.IncrStock(ctx, &rpcproduct.IncrStockReq{
		Id:   preOrder.ProductId,
		Incr: 1,
	})
	if err != nil {
		return fmt.Errorf("increase product stock failed: %w", err)
	}

	// 更新预占状态
	if err := tx.Model(&model.PreOrder{}).Where("id = ?", msg.TempID).Update("status", "cancelled").Error; err != nil {
		return fmt.Errorf("update pre order failed: %w", err)
	}

	// 更新订单状态
	if err := tx.Model(&model.Order{}).Where("pre_order_id = ?", msg.TempID).Update("order_state", model.OrderStateCanceled).Error; err != nil {
		return fmt.Errorf("update order failed: %w", err)
	}

	// 回滚redis
	tempKey := redis.GetSeckillTempKey(msg.TempID)
	if err = redis.RedisClient.Del(ctx, tempKey).Err(); err != nil {
		return fmt.Errorf("delete temp key failed: %w", err)
	}

	if err = redis.RedisClient.Incr(ctx, redis.GetProductStockKey(msg.ProductID)).Err(); err != nil {
		return fmt.Errorf("increase product stock failed: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %w", err)
	}

	return nil
}
