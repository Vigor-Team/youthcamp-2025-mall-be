package redis

import (
	"context"
	"fmt"
	"time"
)

const (
	BEGIN_TIMESTAMP = 1720876200
	COUNT_BITS      = 32
)

func NextId(ctx context.Context, keyPrefix string) (uint32, error) {
	now := time.Now()
	timeStamp := now.Unix() - BEGIN_TIMESTAMP

	date := now.Format("2006:01:02")

	key := fmt.Sprintf("icr:%s:%s", keyPrefix, date)
	count, err := RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment redis key: %v", err)
	}

	id := (timeStamp << COUNT_BITS) | count
	return uint32(id), nil
}
