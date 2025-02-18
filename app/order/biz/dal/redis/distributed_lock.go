package redis

import (
	"context"
	"log"
	"time"
)

func ReleaseLock(ctx context.Context, lockKey string) error {
	luaScript := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`
	_, err := RedisClient.Eval(ctx, luaScript, []string{lockKey}, "mall_lock").Result()
	return err
}

func TryLock(ctx context.Context, lockKey string, ttl time.Duration) bool {
	success, err := RedisClient.SetNX(ctx, lockKey, "mall_lock", ttl).Result()
	if err != nil {
		log.Println("Error acquiring lock:", err)
		return false
	}
	return success
}
