package usersctrl

import (
	"github.com/gofiber/fiber"
)

func Index(c *fiber.Ctx) {
	//fmt.Printf(c.Params("*"))
	c.Send("Index Users")
}

func Show(c *fiber.Ctx) {

	c.Send("Show User")
}

func Create(c *fiber.Ctx) {

	c.Send("Create User")
}

func Store(c *fiber.Ctx) {
	c.Send("Store Users")
}

func Edit(c *fiber.Ctx) {

	c.Send("Edit Users")
}

func Update(c *fiber.Ctx) {

	c.Send("Update Users")
}

func Destroy(c *fiber.Ctx) {

	c.Send("Destroy Users")
}
