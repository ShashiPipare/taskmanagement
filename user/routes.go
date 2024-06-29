package user

import "github.com/gofiber/fiber/v2"

func Route(router fiber.Router) {
	grp := router.Group("/user")
	grp.Post("/signup", signup)
}
