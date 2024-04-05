package http

import (
	"fmt"
	"log"

	"github.com/banggibima/go-fiber-redis/config"
	userapplication "github.com/banggibima/go-fiber-redis/internal/application/user"
	userinfrasctructure "github.com/banggibima/go-fiber-redis/internal/infrasctructure/user"
	"github.com/banggibima/go-fiber-redis/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Fiber  *fiber.App
	Redis  *redis.Client
	Config *config.Config
}

func AppServer(s *Server) {
	userMemory := userinfrasctructure.NewUserMemory(s.Redis)
	userService := userapplication.NewUserService(userMemory)
	userHandler := handler.NewUserHandler(userService)

	r := Router{
		App:         s.Fiber,
		UserHandler: userHandler,
	}

	r.AppRouter()

	port := s.Config.App.Port

	err := s.Fiber.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
