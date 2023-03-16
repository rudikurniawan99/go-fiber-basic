package product_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
	"github.com/rudikurniawan99/go-fiber-basic/src/helper/validator"
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

func CreateProduct(c *fiber.Ctx) error {
	product := &dtos.ProductDTO{}

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "body parser failed",
		})
	}

	if err := validator.Validate(product); err != nil {
		return c.Status(400).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "validation failed",
		})
	}

	return c.Status(201).JSON(res.JsonResponse{
		Success: true,
		Data:    product,
	})
}
