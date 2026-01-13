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

type BucketRequest struct {
	Bucket  string `json:"bucket"`
	Region  string `json:"region"`
	Key     string `json:"key"`
	Profile string `json:"profile"`
}

func BucketExists(c *fiber.Ctx) error {
	var req BucketRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Profile == "" {
		req.Profile = "default"
	}

	if req.Bucket == "" || req.Region == "" || req.Key == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(req.Region),
		config.WithSharedConfigProfile(req.Profile),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "aws configuration failed")
	}

	client := s3.NewFromConfig(cfg)

	if _, err = client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(req.Bucket),
	}); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "bucket not accessible")
	}

	if _, err = client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(req.Bucket),
		Key:    aws.String(req.Key),
	}); err != nil {
		var notFound *s3types.NotFound
		if errors.As(err, &notFound) {
			return c.JSON(fiber.Map{
				"success": true,
				"status":  "bucket_ok_key_missing",
			})
		}
		return fiber.NewError(fiber.StatusInternalServerError, "object lookup failed")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "backend_ok",
	})
}

func CreateBucket(c *fiber.Ctx) error {
	var req BucketRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Profile == "" {
		req.Profile = "default"
	}

	if req.Bucket == "" || req.Region == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(req.Region),
		config.WithSharedConfigProfile(req.Profile),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "aws configuration failed")
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(req.Bucket),
		CreateBucketConfiguration: &s3types.CreateBucketConfiguration{
			LocationConstraint: s3types.BucketLocationConstraint(req.Region),
		},
	})
	if err != nil {
		var exists *s3types.BucketAlreadyOwnedByYou
		if errors.As(err, &exists) {
			return c.JSON(fiber.Map{
				"success": true,
				"status":  "bucket_already_exists",
			})
		}
		return fiber.NewError(fiber.StatusInternalServerError, "bucket creation failed")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "bucket_created",
	})
}

func DeleteBucket(c *fiber.Ctx) error {
	var req BucketRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Profile == "" {
		req.Profile = "default"
	}

	if req.Bucket == "" || req.Region == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(req.Region),
		config.WithSharedConfigProfile(req.Profile),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "aws configuration failed")
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(req.Bucket),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "bucket deletion failed")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "bucket_deleted",
	})
}
