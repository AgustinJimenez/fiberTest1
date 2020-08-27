package routes

import (
	usersctrl "../controllers"
	"github.com/gofiber/fiber"
)

func UsersRoutes(router fiber.Router) {
	usersRouter := router.Group("users")
	usersRouter.Get("/", usersctrl.Index)
	usersRouter.Get("/:id", usersctrl.Show)
	usersRouter.Get("/new", usersctrl.Create)
	usersRouter.Post("/new", usersctrl.Store)
	usersRouter.Get("/edit", usersctrl.Edit)
	usersRouter.Post("/edit", usersctrl.Update)

	//app.("/api/v1/users/edit", usersctrl.destroy)
}
