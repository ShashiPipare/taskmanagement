package data

import "github.com/gofiber/fiber/v2"

func New(c *fiber.Ctx) (a *model) {
	a = &model{
		c,
	}
	return
}
