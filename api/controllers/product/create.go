package product

import "github.com/gofiber/fiber/v2"

func Create(c *fiber.Ctx) error {
	return c.SendString("Created a new product")
}
