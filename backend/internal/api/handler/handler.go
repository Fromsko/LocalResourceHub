package handler

import (
	"LocalResourceHub/backend/internal/api/handler/auth"
	"LocalResourceHub/backend/internal/api/handler/repository"
	"LocalResourceHub/backend/internal/api/handler/upload"
	database "LocalResourceHub/backend/pkg/db"
	"LocalResourceHub/backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	uploadApi := v1.Group("/upload")
	{
		repo := repository.NewGormFileRepository(database.DBConn)
		uploadService := upload.NewUploadService(repo)

		uploadApi.Post("/", uploadService.UploadFile)
		uploadApi.Get("/:id", uploadService.GetFile)
		uploadApi.Delete("/:id", uploadService.DeleteFile)
	}

	authApi := v1.Group("/auth")
	{
		authService := auth.NewAuthService(database.DBConn)
		authApi.Post("/login", authService.Login)
		authApi.Post("/register", authService.Register)
		authApi.Post("/logout", authService.Logout)
	}

	filesApi := v1.Group("/files")
	{
		repo := repository.NewGormFileRepository(database.DBConn)

		filesApi.Get("/", func(c *fiber.Ctx) error {
			page := c.QueryInt("page", 1)
			pageSize := c.QueryInt("pageSize", 1000)
			fileType := c.Query("type", "all")

			files, total, err := repo.List(page, pageSize, fileType)
			if err != nil {
				return utils.JsonErrorResp(c, 500)
			}

			return utils.JsonSuccessResp(c, fiber.Map{
				"files": files,
				"total": total,
			})
		})
	}
}
