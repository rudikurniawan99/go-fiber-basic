package response

import "github.com/gofiber/fiber/v2"

type JsonResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"error"`
	Data    fiber.Map `json:"data"`
}
