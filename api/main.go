package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"weatherApp/api/database"
)

func main() {

	app := fiber.New(fiber.Config{})

	database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("merhaba")
	})

	log.Fatal(app.Listen(":8080"))
}
