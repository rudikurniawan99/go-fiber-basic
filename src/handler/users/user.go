package user_handler

import "github.com/gofiber/fiber/v2"

func GetUserHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "success to get user",
	})
}
