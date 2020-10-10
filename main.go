package main

import (
	"flag"
	"FoundationHelper_KnightHacks2020/app"
	"FoundationHelper_KnightHacks2020/config"
	"FoundationHelper_KnightHacks2020/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

var (
	port = flag.String("port", ":3001", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	Serve()
}

// Serve function starts the web server
func Serve() {
	// Create web server instance
	Boot()

	// Middleware
	app.App.Use(recover.New())
	app.App.Use(logger.New())

	// Load in the http routes
	routes.Load()

	app.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Handle not founds
	//app.Use(handlers.NotFound)

	log.Fatal(app.App.Listen(*port))
}

// Boot function loads intial settings
func Boot() {
	config.BootApp()
}
