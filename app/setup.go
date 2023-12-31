package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goellavish10/skillbee-assignment/config"
	"github.com/goellavish10/skillbee-assignment/lib"
	"github.com/goellavish10/skillbee-assignment/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {

	err := config.LoadENV()
	if err != nil {
		return err
	}
	// Get the current working directory
	currentDir, err := os.Getwd()
	fmt.Println(currentDir)
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	// Static Directory
	// Construct the absolute path to the "resources" directory
	resourcesDir := filepath.Join(currentDir, "resources")
	app.Static("/static/resources", resourcesDir)

	lib.GenerateStaticPages()
	routes.SetupRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)
	return nil
}
