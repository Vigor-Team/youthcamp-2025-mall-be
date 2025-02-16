local product_key = KEYS[1]
local temp_key = KEYS[2]

local user_id = ARGV[1]
local product_id = ARGV[2]
local expire_seconds = ARGV[3]


if redis.call('SISMEMBER', product_key, user_id) == 1 then
    return {err='DUPLICATE_USER'}
end

local stock = redis.call('GET', product_key)
if not stock or tonumber(stock) <= 0 then
    return {err='OUT_OF_STOCK'}
end

local remaining_stock = redis.call('DECRBY', product_key, 1)

if remaining_stock < 0 then
    redis.call('INCRBY', product_key, 1)
    return {err='OUT_OF_STOCK'}
end

redis.call('HSET', temp_key,
        'user_id', user_id,
        'product_id', product_id,
        'status', 'pre_held'
)
redis.call('EXPIRE', temp_key, expire_seconds)

return remaining_stock
