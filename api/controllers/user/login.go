package user

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.Status(200).SendString("user has been logged in")
}
