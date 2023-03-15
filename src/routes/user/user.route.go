package routes_user

import (
	"github.com/gofiber/fiber/v2"
	user_handler "github.com/rudikurniawan99/go-fiber-basic/src/handler/users"
)

type UserRouter interface {
	Mount(group fiber.Router)
}

func Mount(group fiber.Router) {
	group.Get("/", user_handler.GetUserHandler)
	group.Get("/:id", user_handler.GetUserByIdHandler)
	group.Post("/", user_handler.CreateUserHandler)
}
