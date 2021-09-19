package user

import "github.com/gofiber/fiber/v2"

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("This user has been deleted")

}
