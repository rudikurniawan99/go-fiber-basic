package product_route

import (
	"github.com/gofiber/fiber/v2"
	product_handler "github.com/rudikurniawan99/go-fiber-basic/src/handler/products"
)

type ProductRoute interface {
	Mount(group fiber.Router)
}

func Mount(group fiber.Router) {
	group.Get("/", product_handler.GetProduct)
	group.Post("/", product_handler.CreateProduct)
}
