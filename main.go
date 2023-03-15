package main

import (
	"github.com/gofiber/fiber/v2"
	route "github.com/rudikurniawan99/go-fiber-basic/src/routes"
)

func main() {
	app := fiber.New()

	route.Init(app)
	app.Listen(":8080")
}
