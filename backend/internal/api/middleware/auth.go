package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupMiddlewares sets up CORS and logger middleware
func SetupMiddlewares(app *fiber.App) {
	// Use CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Get("Authorization")
		if authToken == "" {
			// TODO: will be implemented
			// return c.Status(http.StatusUnauthorized).JSON(ErrorResponse(http.StatusUnauthorized, "未授权"))
		}
		return c.Next()
	}
}
