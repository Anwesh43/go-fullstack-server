package server

import (
	"github.com/gofiber/fiber/v2"
	"fullstackserver/handlers"
)

func Start() {
	app := fiber.New()
	app.Get("/hello", handlers.HelloWorldHandler)
	app.Get("/hello/:value", handlers.HelloParamHandler)
	app.Listen(":3000")
}