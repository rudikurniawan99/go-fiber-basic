package main

import (
	"github.com/gofiber/fiber/v2"
	user_handler "github.com/rudikurniawan99/go-fiber-basic/src/handler/users"
)

func main() {
	app := fiber.New()

	users := app.Group("users")
	users.Get("/", user_handler.GetUserHandler)

	app.Listen(":8080")
}
