package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app fiber.Router) {
	api := app.Group("/user")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("merhaba")
	})
}
