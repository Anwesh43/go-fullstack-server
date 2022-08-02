package server

import (
	"github.com/gofiber/fiber/v2"
	"fullstackserver/handlers"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	app := fiber.New()
	app.Get("/hello", handlers.HelloWorldHandler)
	app.Get("/hello/:value", handlers.HelloParamHandler)
	app.Get("/migrate/:tableName", func(c *fiber.Ctx) error {
		handlers.Migrate(c, db)
		return nil 
	})
	app.Get("/createbook/:name/:author/:price", func(c *fiber.Ctx) error {
		handlers.CreateBook(c, db)
		return nil 
	})
	app.Listen(":3000")
}