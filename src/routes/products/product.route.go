package product_route

import "github.com/gofiber/fiber/v2"

type ProductRoute interface {
	Mount(group fiber.Router)
}

func Mount(group fiber.Router) {
	// group.Get("/", )
}
