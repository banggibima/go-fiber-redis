package http

import (
	"github.com/banggibima/go-fiber-redis/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App         *fiber.App
	UserHandler *handler.UserHandler
}

func (r *Router) AppRouter() {
	api := r.App.Group("/api")

	users := api.Group("/users")
	users.Get("/", r.UserHandler.GetAllUsers)
	users.Get("/:id", r.UserHandler.GetUserByID)
	users.Post("/", r.UserHandler.CreateUser)
	users.Put("/:id", r.UserHandler.UpdateUser)
	users.Delete("/:id", r.UserHandler.DeleteUser)
}
