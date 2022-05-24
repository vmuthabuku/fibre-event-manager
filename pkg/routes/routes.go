package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmuthabuku/fibre-event-manager/pkg/controllers"
	"github.com/vmuthabuku/fibre-event-manager/pkg/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
	app.Post("/api/admin/login", controllers.Login)
	app.Get("/api/admin/user", controllers.User)
	app.Post("/api/admin/logout", controllers.Logout)

	adminAuthenticated := app.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/api/admin/user", controllers.User)
	adminAuthenticated.Post("/api/admin/login", controllers.Login)

}
