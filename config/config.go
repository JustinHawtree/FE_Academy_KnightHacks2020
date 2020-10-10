package config

import (
	"FoundationHelper_KnightHacks2020/app"
	"github.com/gofiber/fiber/v2"
)

// BootApp creates the web server instance
func BootApp() {
	// Create fiber app
	app.App = fiber.New(fiber.Config{
		Prefork: false,
	})
}
