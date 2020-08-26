package routes

import (
	usersctrl "../controllers"
	"github.com/gofiber/fiber"
)

func RoutesSetup(app *fiber.App) {
	app.Get("/api/v1/users", usersctrl.Index)
	app.Get("/api/v1/users/:id", usersctrl.Show)
	app.Get("/api/v1/users/new", usersctrl.Create)
	app.Post("/api/v1/users/new", usersctrl.Store)
	app.Get("/api/v1/users/edit", usersctrl.Edit)
	app.Post("/api/v1/users/edit", usersctrl.Update)
	//app.("/api/v1/users/edit", usersctrl.destroy)
}
