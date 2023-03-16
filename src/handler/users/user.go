package user_handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
	"github.com/rudikurniawan99/go-fiber-basic/src/models"
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
	validate := validator.New()
	user := dtos.UserDTO{}

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(res.JsonResponse{
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
	}

	if err := validate.Struct(&user); err != nil {
		return c.JSON(res.JsonResponse{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(res.JsonResponse{
		Success: true,
		Data:    user,
	})
}

func GetUserByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.JSON(res.JsonResponse{
			Success: false,
			Error:   "data not found",
			Data:    nil,
		})
	}

	return c.JSON(res.JsonResponse{
		Success: true,
		Data: models.UserModel{
			ID:       id,
			Name:     "K Kurniawawn",
			Email:    "kurniawan401@gmail.com",
			Password: "kurniawan",
		},
	})
}
