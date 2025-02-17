package auth

import (
	"LocalResourceHub/backend/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) model.AuthHandler {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) Login(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthService) Register(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthService) Logout(ctx *fiber.Ctx) error {
	return nil
}
