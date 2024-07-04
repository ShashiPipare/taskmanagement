package user

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	grp := router.Group("/user")
	grp.Post("/signUp", signUp)
	grp.Post("/login", login)
	grp.Post("/logout", logout)
}
