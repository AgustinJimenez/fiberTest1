package main

import (
	"./app/routes"
	"./app/setup"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
)

func main() {
	app := fiber.New()
	app.Use(helmet.New())
	// 80 requests per 60 seconds max
	cfg := limiter.Config{
		Timeout: 60,
		Max:     80,
	}
	setup.InitDatabase()
	routes.RoutesSetup(app)

	app.Use(limiter.New(cfg))

	app.Listen(3000)
}
