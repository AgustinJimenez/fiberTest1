package routes

import (
	"github.com/gofiber/fiber"
)

func RoutesSetup(app *fiber.App) {
	router := app.Group("/api/v1")
	UsersRoutes(router)
}
