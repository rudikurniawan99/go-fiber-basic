package product_handler

import (
	"github.com/gofiber/fiber/v2"
	res "github.com/rudikurniawan99/go-fiber-basic/src/response"
)

func GetProduct(c *fiber.Ctx) error {
	return c.JSON(res.JsonResponse{
		Success: true,
		Data: fiber.Map{
			"name":  "Hp Bagus",
			"price": 3000000,
		},
	})
}
