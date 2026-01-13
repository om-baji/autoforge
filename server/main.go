package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/om-baji/Auto-Forge/service"
)

func main() {

	app := fiber.New()

	app.Use("/health", service.Health)

	app.Post("/bucket", service.BucketExists)
	app.Delete("/bucket", service.DeleteBucket)
	app.Post("/bucket/new", service.CreateBucket)

	log.Fatal(app.Listen(":8000"))
}
