package main

import (
	"log"
	"simple-rest-api/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configuration := config.Load()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  true,
			"message": "200 OK",
		})
	})

	log.Fatal(app.Listen(configuration.Host.Host + ":" + configuration.Host.Port))
}
