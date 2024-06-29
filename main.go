package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Starting server...")
	port := ":7000"
	app := fiber.New()
	fmt.Println("server started at port", port)
	app.Listen(port)
}
