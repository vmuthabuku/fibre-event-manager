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
	app.Post("/api/ambassador/register", controllers.Register)
	app.Post("/api/ambassador/login", controllers.Login)

	adminAuthenticated := app.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/api/admin/user", controllers.User)
	adminAuthenticated.Post("/api/admin/login", controllers.Login)
	adminAuthenticated.Put("/api/admin/info", controllers.UpdateUser)
	adminAuthenticated.Put("/api/admin/password", controllers.UpdatePassword)
	adminAuthenticated.Get("/api/admin/ambassadors", controllers.Ambassador)
	adminAuthenticated.Get("/api/admin/products", controllers.Products)
	adminAuthenticated.Post("/api/admin/products", controllers.CreateProducts)
	adminAuthenticated.Get("/api/admin/products/:id", controllers.GetProduct)
	adminAuthenticated.Put("/api/admin/products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("/api/admin/products/:id", controllers.DeleteProduct)
	adminAuthenticated.Get("/api/admin/users/:id/links", controllers.Link)
	adminAuthenticated.Get("/api/admin/orders", controllers.Order)

	// ambassador auth middleware application
	ambassadorAuthenticated := app.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("/api/ambassador/user", controllers.User)
	ambassadorAuthenticated.Post("/api/ambassador/login", controllers.Login)
	ambassadorAuthenticated.Put("/api/ambassador/info", controllers.UpdateUser)
	ambassadorAuthenticated.Put("/api/ambassador/password", controllers.UpdatePassword)

}
