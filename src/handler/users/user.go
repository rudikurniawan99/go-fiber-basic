package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudikurniawan99/go-fiber-basic/src/dtos"
	bcrypt_helper "github.com/rudikurniawan99/go-fiber-basic/src/helper/bcrypt"
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
		return c.Status(fiber.StatusBadRequest).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "body parser error",
		})
	}

	if err := validator.ValidateUserPassword(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "password should have 8 character,at least 1 uppercase, 1 lowercase, 1 special character",
		})
	}

	hashPassword, err := bcrypt_helper.GeneratePassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse{
			Success: false,
			Error:   err.Error(),
			Message: "failed to hash password",
		})
	}
	user.Password = string(hashPassword)

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
