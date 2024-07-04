package tasks

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	grp := router.Group("/tasks")
	grp.Post("/tasks", add)
	grp.Put("/tasks", update)
	grp.Get("/tasks/:id", getByID)
	grp.Get("/tasks", getAllTasks)
	grp.Delete("/tasks/:id", delete)
}
