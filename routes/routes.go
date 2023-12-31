package routes

import (
	"github.com/goellavish10/skillbee-assignment/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/", handlers.RunRenderAgain)
	app.Get("/static/:pageId", handlers.RenderDynamicPage)
}
