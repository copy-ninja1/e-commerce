package user

import "github.com/gofiber/fiber/v2"

func UpdateUser(c *fiber.Ctx) error {
	return c.Status(201).SendString("user updated")
}
