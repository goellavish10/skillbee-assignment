package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	currendtDir, err := os.Getwd()
	if err != nil {
		return c.Status(500).SendString("Server Error")
	}
	return c.SendFile(currendtDir + "/views/main.html")
}

func RenderDynamicPage(c *fiber.Ctx) error {
	currendtDir, err := os.Getwd()
	if err != nil {
		return c.Status(500).SendString("Server Error")
	}
	return c.SendFile(currendtDir + fmt.Sprintf("/dist/page-%s.html", c.Params("pageId")))
}
