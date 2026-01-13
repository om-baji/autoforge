package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type HealthResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func Health(c *fiber.Ctx) error {
	return c.JSON(HealthResponse{
		Status: "running!",
		Time:   time.Now(),
	})
}
