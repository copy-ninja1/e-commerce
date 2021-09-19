package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/e-commerce/config/routes"
)

func main() {
	app := fiber.New()
	routes.V1(app, "/api")
	log.Fatal(app.Listen(":3000"))
}
