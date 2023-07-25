package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"TITLE":    "TEST PAGE",
		"BODYTEXT": "Hello World!",
	})
}

func RenderDynamicPage(c *fiber.Ctx) error {
	return c.SendString("OK")
}
