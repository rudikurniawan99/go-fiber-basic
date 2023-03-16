package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
	validator "github.com/rudikurniawan99/go-fiber-basic/src/helper/validator"
	"github.com/rudikurniawan99/go-fiber-basic/src/models"
	res "github.com/rudikurniawan99/go-fiber-basic/src/response"
)

func GetUserHandler(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"name": "Rudi Kurniawan",
			"age":  24,
		},
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	user := &dtos.UserDTO{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "body parser error",
		})
	}

	if err := validator.Validate(user); err != nil {
		return c.Status(400).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "validator fail",
		})
	}

	return c.Status(200).JSON(res.JsonResponse{
		Success: true,
		Data:    user,
	})
}

func GetUserByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(404).JSON(res.ErrorResponse{
			Success: false,
			Error:   "data not found",
		})
	}

	return c.Status(201).JSON(res.JsonResponse{
		Success: true,
		Data: models.UserModel{
			ID:       id,
			Name:     "K Kurniawawn",
			Email:    "kurniawan401@gmail.com",
			Password: "kurniawan",
		},
	})
}
