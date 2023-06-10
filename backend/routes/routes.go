package routes

import (
	"azril/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Get("/reads", controllers.Reads)
	app.Get("/read/:id", controllers.Read)
	app.Post("/create", controllers.Create)
	app.Put("/update/:id", controllers.Update)
	app.Delete("/delete/:id", controllers.Delete)
}
