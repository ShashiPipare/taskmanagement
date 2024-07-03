package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"main.go/config"
	"main.go/connection"
	"main.go/router"
)

func main() {
	fmt.Println("Starting server...")
	conf := config.Init()
	port := conf.Port
	connection.Init(conf)
	connection.ConnectDB()
	app := fiber.New()
	router.Configure(app)
	fmt.Println("server started at port", port)
	app.Listen(port)
}
