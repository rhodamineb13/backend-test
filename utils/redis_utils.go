package utils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func DeleteRedisKeyMatchingPattern(ctx context.Context, pattern string, rc *redis.Client) error {
	iter := rc.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		err := rc.Unlink(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
