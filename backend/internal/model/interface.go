package model

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// Uploader
	Uploader interface {
		FileValidator
		UploadFile(ctx *fiber.Ctx) error
		GetFile(ctx *fiber.Ctx) error
		DeleteFile(ctx *fiber.Ctx) error
		GetAllowedFileTypes() map[string]bool
	}

	// AuthHandler
	AuthHandler interface {
		Login(ctx *fiber.Ctx) error
		Register(ctx *fiber.Ctx) error
		Logout(ctx *fiber.Ctx) error
	}
)
