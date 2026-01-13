package service

import (
	"github.com/gofiber/fiber/v2"
)

type Ec2Request struct {
	Region   string `json:"region"`
	Instance string `json:"instance"`
	Ami      string `json:"ami"`
	Gateway  bool   `json:"gateway"`
	Ssh      string `json:"ssh"`
}

func ValidateEC2(ctx *fiber.Ctx) error {
	var req Ec2Request

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if req.Region == "" ||
		req.Instance == "" ||
		req.Ami == "" ||
		req.Ssh == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	return ctx.Next()
}
