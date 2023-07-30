package handlers

import (
	"fmt"
	"os"

	"github.com/goellavish10/skillbee-assignment/interfaces"
	"github.com/goellavish10/skillbee-assignment/lib"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	currendtDir, err := os.Getwd()
	if err != nil {
		return c.Status(500).SendString("Server Error")
	}
	return c.SendFile(currendtDir + "/views/main.html")
}

func RunRenderAgain(c *fiber.Ctx) error {
	bodyData := new(interfaces.FormData)
	if err := c.BodyParser(bodyData); err != nil {
		fmt.Println("Error while parsing form data: ")
		return err
	}
	if bodyData.NumberOfPages < 1 || bodyData == nil {
		return c.Status(400).SendString("Invalid number of pages")
	}
	no := fmt.Sprintf("%d", bodyData.NumberOfPages)
	os.Setenv("PAGES", no)
	lib.GenerateStaticPages()
	return c.SendString("🎊Static site generated succesfully!")

}

func RenderDynamicPage(c *fiber.Ctx) error {
	currendtDir, err := os.Getwd()
	if err != nil {
		return c.Status(500).SendString("Server Error")
	}
	_, err = os.Stat(currendtDir + fmt.Sprintf("/dist/page-%s.html", c.Params("pageId")))
	if err != nil {
		return c.Status(404).SendString("Page not found")
	}
	return c.SendFile(currendtDir + fmt.Sprintf("/dist/page-%s.html", c.Params("pageId")))
}
