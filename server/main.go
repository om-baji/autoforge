package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/om-baji/Auto-Forge/service"
)

func main() {

	app := fiber.New()

	app.Use("/health", service.Health)

	app.Post("/terraform", service.CheckTF)

	log.Fatal(app.Listen(":3000"))
}
