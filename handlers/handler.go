package handlers 

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"fullstackserver/db"
	"gorm.io/gorm"
	"fullstackserver/models"
	"strconv"
)

func HelloWorldHandler(c *fiber.Ctx) error {
	c.SendString("Hello World")
	return nil 
}

func HelloParamHandler(c *fiber.Ctx) error {
	c.SendString(fmt.Sprintf("Hello  %s", c.Params("value")))
	return nil 
}

func Migrate(c *fiber.Ctx, database *gorm.DB) error {
	db.Migrate(database, c.Params("tableName"))
	c.SendString(fmt.Sprintf("Migrated %s", c.Params("tableName")))
	return nil 
}

func CreateBook(c *fiber.Ctx, database *gorm.DB) error {
	price, err := strconv.ParseUint(c.Params("price"), 10, 64)
	if err != nil {
		price = 0
	}
	db.Create(database, &models.Book{
		Name: c.Params("name"),
		Author: c.Params("author"),
		Price: uint(price),
	})
	c.SendString("successfully created book")
	return nil 
}

func GetBooks(c *fiber.Ctx, database *gorm.DB) error {
	books := make([]models.Book, 0)
	db.QueryAll(database, &books)
	c.JSON(books)
	return nil 
}

// fun CreateAuthor(c *fiber.Ctx, database *gorm.DB) error {
// 	db.Create(database, &models.Author{
// 		name: c.Params("name"),
// 		country: c.Params("country"),
// 	})
// 	c.SendString("successfully created book")
// 	return nil 
// }