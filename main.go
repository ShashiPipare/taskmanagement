package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"main.go/connection"
)

func main() {
	fmt.Println("Starting server...")
	port := ":7000"
	connection.Init()
	connection.ConnectDB()
	app := fiber.New()
	fmt.Println("server started at port", port)
	app.Listen(port)
}
