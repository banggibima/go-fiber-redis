package main

import (
	"github.com/banggibima/go-fiber-redis/config"
	"github.com/banggibima/go-fiber-redis/internal/interface/http"
	pkgfiber "github.com/banggibima/go-fiber-redis/pkg/fiber"
	pkgredis "github.com/banggibima/go-fiber-redis/pkg/redis"
)

func main() {
	cfg, err := config.Conf()
	if err != nil {
		panic(err)
	}

	fiber, err := pkgfiber.FiberInit()
	if err != nil {
		panic(err)
	}

	redis := pkgredis.RedisInit(cfg)

	if err := pkgredis.RedisConnect(redis); err != nil {
		panic(err)
	}

	http.AppServer(&http.Server{
		Fiber:  fiber,
		Config: cfg,
		Redis:  redis,
	})
}
