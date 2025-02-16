package script

import (
	"sync"
)

var (
	scriptOnce       sync.Once
	preSeckillScript string
)

func GetPreSeckillScript() (string, error) {
	//var err error
	//scriptOnce.Do(func() {
	//	prefix := "redis/script"
	//	scriptFileRelPath := filepath.Join(prefix, "pre_seckill.lua")
	//	content, err := os.ReadFile(scriptFileRelPath)
	//	if err != nil {
	//		return
	//	}
	//	preSeckillScript = string(content)
	//})
	//return preSeckillScript, err
	return `
local product_stock_key = KEYS[1]
local product_order_key = KEYS[2]

local user_id = ARGV[1]
local product_id = ARGV[2]
local pre_order_id = ARGV[3]
local expire_seconds = ARGV[4]
  
if redis.call('EXISTS', product_order_key) == 1 then
    return {err='DUPLICATE_USER'}
end

local stock = redis.call('GET', product_stock_key)
if not stock or tonumber(stock) <= 0 then
    return {err='OUT_OF_STOCK'}
end

local remaining_stock = redis.call('DECRBY', product_stock_key, 1)

if remaining_stock < 0 then
    redis.call('INCRBY', product_stock_key, 1)
    return {err='OUT_OF_STOCK'}
end

redis.call('HSET', product_order_key,
        'user_id', user_id,
        'product_id', product_id,
        'pre_order_id', pre_order_id,
        'status', 'pre_held'
)

return remaining_stock

`, nil
}
