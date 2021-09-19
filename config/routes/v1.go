package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/e-commerce/api/controllers/user"
)

func V1(app *fiber.App, groupPath string) {
	api := app.Group(groupPath)
	v1 := api.Group("/v1")

	v1.Post("/user/signup", user.Signup)
	v1.Post("/user/login", user.Login)
	v1.Get("/user/find", user.FindUser)
	v1.Patch("/user/update", user.UpdateUser)
	v1.Delete("/user/delete", user.DeleteUser)

}
