package functions

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.SendString("hello from homepage")
}
