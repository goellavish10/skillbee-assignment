package handlers

import "github.com/gofiber/fiber/v2"

func RenderDynamicPage(c *fiber.Ctx) error {
	return c.SendString("OK")
}
