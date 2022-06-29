package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"weatherApp/api/handlers"
	"weatherApp/api/utils"
	"weatherApp/repository"
	"weatherApp/services"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	repo := repository.NewMysqlRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewHandler(service)

	api := app.Group("/api")
	api.Post("/user", handler.Create)

	log.Fatal(app.Listen(":8080"))
}
