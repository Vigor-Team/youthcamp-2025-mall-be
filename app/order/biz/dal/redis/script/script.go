package script

func GetPreSeckillScript() string {
	return `
local product_stock_key = KEYS[1]
local product_order_key = KEYS[2]
local pre_order_key = KEYS[3]

local user_id = ARGV[1]
local product_id = ARGV[2]

if redis.call('SISMEMBER', product_order_key, user_id) == 1 then
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

redis.call('SADD', product_order_key, user_id)
redis.call('HSET', pre_order_key, 
    "user_id", user_id,
    "product_id", product_id
)

return remaining_stock
`
}
