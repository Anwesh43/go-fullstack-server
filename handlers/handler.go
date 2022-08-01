package handlers 

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func HelloWorldHandler(c *fiber.Ctx) error {
	c.SendString("Hello World")
	return nil 
}

func HelloParamHandler(c *fiber.Ctx) error {
	c.SendString(fmt.Sprintf("Hello  %s", c.Params("value")))
	return nil 
}