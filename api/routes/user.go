package routes

import (
	"github.com/gofiber/fiber/v2"
	"weatherApp/api/utils"
)

func UserRoutes(app fiber.Router) {
	api := app.Group("/user")

	api.Get("/", func(c *fiber.Ctx) error {
		return utils.SuccessPresenter(c, "merhaba", nil)
	})
}
