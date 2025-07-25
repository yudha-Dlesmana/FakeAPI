package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yudha-Dlesmana/fakeAPI/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", handlers.WelcomeHandler)
	v1.Get("/fakeAPI", handlers.FakeApiHandler)
}
