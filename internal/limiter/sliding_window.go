package limiter

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewSlidingWindowLimiter(client *redis.Client, limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (l *SlidingWindowLimiter) Allow(key string) (bool, error) {
	ctx := context.Background()
	now := time.Now().UnixMilli()
	windowStart := now - l.window.Milliseconds()

	pipe := l.client.Pipeline()

	// Remove old entries
	pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(windowStart, 10))

	// Count requests in window
	countCmd := pipe.ZCard(ctx, key)

	// Add current request
	pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})

	// Set expiry
	pipe.Expire(ctx, key, l.window)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	return countCmd.Val() < int64(l.limit), nil
}


