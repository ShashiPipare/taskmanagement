package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	fmt.Println("11")

	grp := router.Group("/user")
	grp.Post("/signUp", signUp)
}
