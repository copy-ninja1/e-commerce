package product

import "github.com/gofiber/fiber/v2"

func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete user")
}
