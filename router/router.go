package router

import (
	"github.com/gofiber/fiber/v2"
	"main.go/tasks"
	"main.go/user"
)

func Configure(app *fiber.App) {
	api := app.Group("/api")
	tasks.Route(api)
	user.Route(api)
}
