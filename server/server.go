package server

import (
	"github.com/gofiber/fiber/v2"
	"fullstackserver/handlers"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(db *gorm.DB) {
	app := fiber.New()
	app.Use(cors.New())
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
	app.Get("/books", func (c *fiber.Ctx) error {
		return handlers.GetBooks(c, db)
	})

	app.Listen(":3000")
}