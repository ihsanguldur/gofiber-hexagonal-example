package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"weatherApp/api/database"
	"weatherApp/api/routes"
	"weatherApp/api/utils"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	database.Connect()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
