package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	route "github.com/rudikurniawan99/go-fiber-basic/src/routes"
)

func main() {
	app := fiber.New()

	route.Init(app)
	app.Listen(":" + os.Getenv("PORT"))
}
