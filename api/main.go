package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"weatherApp/api/database"
	"weatherApp/api/routes"
)

func main() {

	app := fiber.New(fiber.Config{})

	database.Connect()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
