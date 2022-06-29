package utils

import "github.com/gofiber/fiber/v2"

func SuccessPresenter(c *fiber.Ctx, message string, data interface{}) error {
	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"success": true,
			"message": message,
			"data":    data,
		})
}

func ErrorPresenter(c *fiber.Ctx, message string, code int) error {
	return c.
		Status(code).
		JSON(fiber.Map{
			"success": false,
			"message": message,
			"data":    nil,
		})
}
