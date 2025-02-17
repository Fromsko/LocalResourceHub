package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewCORSMiddleware returns a new CORS middleware instance
func NewCORSMiddleware() fiber.Handler {
	config := cors.Config{
		AllowOrigins:     "http://localhost:3000,http://wails.localhost:34115",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}

	return cors.New(config)
}
