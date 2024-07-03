package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"main.go/tasks"
	"main.go/user"
)

func Configure(app *fiber.App) {
	fmt.Println("1")
	api := app.Group("/api")
	tasks.Route(api)
	user.Route(api)
}
