package redis

import (
	"fmt"

	"github.com/banggibima/go-fiber-redis/config"
	"github.com/redis/go-redis/v9"
)

func RedisInit(cfg *config.Config) *redis.Client {
	host := cfg.Redis.Host
	port := cfg.Redis.Port
	password := cfg.Redis.Password
	database := cfg.Redis.Database

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       database,
	}

	client := redis.NewClient(options)

	return client
}
