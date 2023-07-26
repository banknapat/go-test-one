package routes

import (
	"github.com/banknapat/go-test-one/src/controllers"
	"github.com/banknapat/go-test-one/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/hello", controllers.Hello)

	app.Post("/multiply", controllers.MultiplyHandler)

	app.Use(middlewares.NotFoundErr())
}
