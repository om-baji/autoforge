package service

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2"
)

type TerraformInit struct {
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	Profile string `json:"profile"`
	Region  string `json:"region"`
}

func CheckTF(ctx *fiber.Ctx) error {
	var payload TerraformInit
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if payload.Profile == "" {
		payload.Profile = "default"
	}

	if payload.Bucket == "" || payload.Key == "" || payload.Region == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(payload.Region),
		config.WithSharedConfigProfile(payload.Profile),
	)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid aws credentials or profile")
	}

	s3Client := s3.NewFromConfig(awsCfg)

	_, err = s3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(payload.Bucket),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "bucket not accessible")
	}

	_, err = s3Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(payload.Bucket),
		Key:    aws.String(payload.Key),
	})
	if err != nil {
		var notFound *s3types.NotFound
		if errors.As(err, &notFound) {
			return ctx.JSON(fiber.Map{
				"success": true,
				"status":  "bucket_ok_key_missing",
			})
		}
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"status":  "backend_ok",
	})
}
