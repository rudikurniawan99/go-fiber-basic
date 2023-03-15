package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
	res "github.com/rudikurniawan99/go-fiber-basic/src/response"
)

func GetUserHandler(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"name": "Rudi Kurniawan",
			"age":  24,
		},
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	user := dtos.UserDTO{}

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(res.JsonResponse{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(res.JsonResponse{
		Success: true,
		Data: fiber.Map{
			"body": &user,
		},
	})
}
