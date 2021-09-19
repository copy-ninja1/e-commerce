package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/e-commerce/api/controllers/product"
	"github.com/copy-ninja1/e-commerce/api/controllers/user"
)

func V1(app *fiber.App, groupPath string) {
	api := app.Group(groupPath)
	v1 := api.Group("v1")
	usersRoutes(v1)
	productsRoutes(v1)
}

func usersRoutes(v fiber.Router) {

	// ██╗   ██╗███████╗███████╗██████╗     ██████╗  ██████╗ ██╗   ██╗████████╗███████╗███████╗
	// ██║   ██║██╔════╝██╔════╝██╔══██╗    ██╔══██╗██╔═══██╗██║   ██║╚══██╔══╝██╔════╝██╔════╝
	// ██║   ██║███████╗█████╗  ██████╔╝    ██████╔╝██║   ██║██║   ██║   ██║   █████╗  ███████╗
	// ██║   ██║╚════██║██╔══╝  ██╔══██╗    ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  ╚════██║
	// ╚██████╔╝███████║███████╗██║  ██║    ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗███████║
	//  ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝    ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚══════╝

	v.Post("/user/signup", user.Signup)
	v.Post("/user/login", user.Login)
	v.Get("/user/find", user.FindUser)
	v.Patch("/user/update", user.UpdateUser)
	v.Delete("/user/delete", user.DeleteUser)
}

func productsRoutes(v fiber.Router) {
	// ██████╗ ██████╗  ██████╗ ██████╗ ██╗   ██╗ ██████╗████████╗    ██████╗  ██████╗ ██╗   ██╗████████╗███████╗███████╗
	// ██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██║   ██║██╔════╝╚══██╔══╝    ██╔══██╗██╔═══██╗██║   ██║╚══██╔══╝██╔════╝██╔════╝
	// ██████╔╝██████╔╝██║   ██║██║  ██║██║   ██║██║        ██║       ██████╔╝██║   ██║██║   ██║   ██║   █████╗  ███████╗
	// ██╔═══╝ ██╔══██╗██║   ██║██║  ██║██║   ██║██║        ██║       ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  ╚════██║
	// ██║     ██║  ██║╚██████╔╝██████╔╝╚██████╔╝╚██████╗   ██║       ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗███████║
	// ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝  ╚═════╝  ╚═════╝   ╚═╝       ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚══════╝

	v.Post("/product/create", product.Create)
	v.Get("/product/find", product.Find)
	v.Patch("/product/update", product.Update)
	v.Delete("/product/delete", product.Delete)
}
