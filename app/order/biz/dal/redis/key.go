package redis

import "fmt"

func GetSeckillTempLockKey(tempId string) string {
	return fmt.Sprintf("seckill:lock:%s", tempId)
}

func GetProductStockKey(productId uint32) string {
	return fmt.Sprintf("product:stock:%d", productId)
}

func GetProductOrderKey(productId uint32) string {
	return fmt.Sprintf("product:order:%d", productId)
}
