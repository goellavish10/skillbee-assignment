package app

import (
	"os"

	"github.com/goellavish10/skillbee-assignment/config"
	"github.com/goellavish10/skillbee-assignment/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func SetupAndRunApp() error {

	err := config.LoadENV()
	if err != nil {
		return err
	}
	// html engine
	engine := html.New("../views", ".html")
	// create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	// Static Directory
	app.Static("/resources", "../resources")

	routes.SetupRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)
	return nil
}
