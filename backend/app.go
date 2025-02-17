package backend

import (
	"context"
	"LocalResourceHub/backend/internal/api/handler"
	"LocalResourceHub/backend/internal/api/middleware"
	"LocalResourceHub/backend/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// App represents the main application structure
type App struct {
	ctx    context.Context
	logger *Logger
}

// Logger manages application logging
type Logger struct{}

// NewLogger creates a new Logger instance
func NewLogger() *Logger {
	return &Logger{}
}

// Configure sets up logger settings
func (l *Logger) Configure() {
	// Configuration logic can be added here if needed
}

// NewApp creates a new application instance
func NewApp() *App {
	return &App{
		logger: NewLogger(),
	}
}

// Startup initializes and starts the application
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Configure()

	database.InitDB("resouse.db")
	app := a.setupFiberApp()
	
	a.setupMiddlewares(app)
	handler.SetupRoutes(app)

	app.Listen("0.0.0.0:3670")
}

// setupFiberApp configures and returns a new Fiber app instance
func (a *App) setupFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024, // Set body limit to 1GB
	})
}

// setupMiddlewares configures application middlewares
func (a *App) setupMiddlewares(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Shanghai",
	}))
	middleware.SetupMiddlewares(app)
}

// Shutdown handles application shutdown
func (a *App) Shutdown(ctx context.Context) {
	a.ctx = ctx
}

// DomReady handles DOM ready event
func (a *App) DomReady(ctx context.Context) {
	a.ctx = ctx
}
