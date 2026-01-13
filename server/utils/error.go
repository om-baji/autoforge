package utils

import "github.com/gofiber/fiber/v2"

func BadRequest(msg string) error {
	if msg == "" {
		msg = "bad request"
	}
	return fiber.NewError(fiber.StatusBadRequest, msg)
}
