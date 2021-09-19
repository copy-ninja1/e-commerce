package user

import "github.com/gofiber/fiber/v2"

func FindUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("user found")
}
