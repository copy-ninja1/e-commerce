package product

import "github.com/gofiber/fiber/v2"

func Update(c *fiber.Ctx) error {
	return c.SendString("Product updated")
}
