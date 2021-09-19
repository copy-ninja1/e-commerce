package product

import "github.com/gofiber/fiber/v2"

func Find(c *fiber.Ctx) error {
	return c.SendString("found product")
}
