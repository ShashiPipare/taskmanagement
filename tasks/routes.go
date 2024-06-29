package tasks

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	grp := router.Group("/tasks")
	grp.Post("/add", add)
}
