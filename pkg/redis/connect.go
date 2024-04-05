package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func RedisConnect(rdb *redis.Client) error {
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
