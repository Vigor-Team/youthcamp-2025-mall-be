package redis

import (
	"context"
	"fmt"
	"time"
)

const BEGIN_TIMESTAMP int64 = 1720876200
const COUNT_BITS uint32 = 16
const MAX_COUNT uint32 = (1 << COUNT_BITS) - 1

func NextId(ctx context.Context, keyPrefix string) (uint32, error) {
	now := time.Now()
	timestamp := (now.Unix() - BEGIN_TIMESTAMP) / 60

	date := now.Format("2006:01:02")
	count, err := RedisClient.Incr(ctx, fmt.Sprintf("icr:%s:%s", keyPrefix, date)).Result()
	if err != nil {
		return 0, err
	}
	countUint32 := uint32(count)

	if countUint32 > MAX_COUNT {
		countUint32 = 0
	}

	id := uint32(timestamp&0xFFF)<<COUNT_BITS | countUint32

	return id, nil
}
