package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/config"
	"github.com/vmuthabuku/fibre-event-manager/pkg/routes"
)

func main() {

	config.Connect()
	config.AutoMigrate()

	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
