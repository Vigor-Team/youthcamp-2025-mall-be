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
		return errors.New("get lock failed")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "release lock failed: %v", err)
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
		return errors.New("add pre order failed")
	}

	// publish delay message
	producer := NewProducer(Client)
	delayMsg := DelayMessage{
		TempID:     msg.TempID,
		CreatedAt:  time.Now().Unix(),
		ExpectedAt: time.Now().Add(10 * time.Minute).Unix(),
	}
	if err = producer.PublishDelay(ctx, delayMsg, 1*time.Second); err != nil {
		return errors.New("publish delay message failed")
	}
	return nil
}

func handleOrder(ctx context.Context, msg OrderMessage) error {
	// get distributed lock
	key := redis.GetSeckillTempLockKey(msg.TempID)
	if success := redis.TryLock(ctx, key, 20*time.Second); !success {
		return errors.New("get lock failed")
	}
	defer func(ctx context.Context, lockKey string) {
		err := redis.ReleaseLock(ctx, lockKey)
		if err != nil {
			klog.CtxErrorf(ctx, "release lock failed: %v", err)
		}
	}(ctx, key)

	// check order exists
	_, err := model.GetOrder(mysql.DB, ctx, msg.UserID, strconv.Itoa(int(msg.OrderId)))
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("order already exists")
	}

	// start transaction
	tx := mysql.DB.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	defer tx.Rollback()

	// update pre order
	if err := tx.Model(&model.PreOrder{}).Where("id = ?", msg.TempID).Update("status", "completed").Update("order_id", msg.OrderId).Error; err != nil {
		return fmt.Errorf("update pre order failed: %v", err)
	}

	// decrease product stock
	_, err = rpc.ProductClient.DecrStock(ctx, &rpcproduct.DecrStockReq{
		Id:   msg.ProductId,
		Decr: 1,
	})
	if err != nil {
		return fmt.Errorf("decrease product stock failed: %v", err)
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
	if err := tx.Create(&itemList).Error; err != nil {
		return err
	}

	// delete redis key
	productOrderKey := redis.GetProductOrderKey(msg.ProductId)
	err = redis.RedisClient.Del(ctx, productOrderKey).Err()
	if err != nil {
		return fmt.Errorf("delete product order failed: %v", err)
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %v", err)
	}
	return nil
}

func handleDelayOrder(ctx context.Context, msg DelayMessage) error {
	// get distributed lock
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
	if err := mysql.DB.Where("id = ?", msg.TempID).First(&preOrder).Error; err != nil {
		return fmt.Errorf("get pre order failed: %v", err)
	}

	// 更新预占状态
	if err := tx.Model(&model.PreOrder{}).Where("id = ?", msg.TempID).Update("status", "cancelled").Error; err != nil {
		return fmt.Errorf("update pre order failed: %v", err)
	}

	// 更新订单状态
	if err := tx.Model(&model.Order{}).Where("pre_order_id = ?", msg.TempID).Update("order_state", model.OrderStateCanceled).Error; err != nil {
		return fmt.Errorf("update order failed: %v", err)
	}

	// 恢复库存
	_, err := rpc.ProductClient.IncrStock(ctx, &rpcproduct.IncrStockReq{
		Id:   preOrder.ProductId,
		Incr: 1,
	})
	if err != nil {
		return fmt.Errorf("increase product stock failed: %v", err)
	}

	// 回滚redis
	productOrderKey := redis.GetProductOrderKey(msg.ProductID)
	if err = redis.RedisClient.SRem(ctx, productOrderKey, msg.UserID).Err(); err != nil {
		return fmt.Errorf("delete user buy record failed: %v", err)
	}

	if err = redis.RedisClient.Incr(ctx, redis.GetProductStockKey(msg.ProductID)).Err(); err != nil {
		return fmt.Errorf("increase product stock failed: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction fail: %v", err)
	}

	return nil
}
