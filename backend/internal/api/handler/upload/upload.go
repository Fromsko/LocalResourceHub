package upload

import (
	"LocalResourceHub/backend/internal/model"
	"LocalResourceHub/backend/pkg/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UploadService struct {
	repo             model.FileRepository
	MaxFileSize      int64
	AllowedFileTypes map[string]bool
}

func NewUploadService(repo model.FileRepository) model.Uploader {
	return &UploadService{
		repo:        repo,
		MaxFileSize: 1024 * 1024 * 1024, // 1GB
		AllowedFileTypes: map[string]bool{
			"application/pdf":    true,
			"application/msword": true, // .doc
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true, // .docx
			"application/vnd.ms-excel": true, // .xls
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true, // .xlsx
			"text/plain":                   true,
			"text/csv":                     true,
			"application/zip":              true,
			"application/x-rar-compressed": true,
			"image/jpeg":                   true,
			"image/png":                    true,
			"image/gif":                    true,
			"application/octet-stream":     true, // Fallback for unknown binary files
		},
	}
}

func (s *UploadService) ValidateFileSize(size int64) error {
	if size > s.MaxFileSize {
		return fmt.Errorf("file size exceeds limit")
	}
	return nil
}
func (s *UploadService) ValidateFileType(contentType string, fileType string) error {
	if !s.AllowedFileTypes[contentType] {
		return fmt.Errorf("unsupported file type")
	}
	return nil
}

func (s *UploadService) GetAllowedFileTypes() map[string]bool {
	return s.AllowedFileTypes
}

func (s *UploadService) UploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_UPLOAD)
	}

	// 验证文件大小
	if file.Size > s.MaxFileSize {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_TOO_LARGE)
	}

	// 验证文件类型
	contentType := file.Header.Get("Content-Type")
	// if !s.AllowedFileTypes[contentType] {
	// 	return utils.JsonErrorResp(ctx, model.ERROR_UNSUPPORTED_TYPE)
	// }

	// 计算文件hash
	hash, err := utils.CalculateFileHash(file)
	if err != nil {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_UPLOAD)
	}

	// 检查文件是否已存在
	existingFile, _ := s.repo.FindByHash(hash)
	if existingFile != nil {
		return utils.JsonSuccessResp(ctx, existingFile)
	}

	// 读取文件内容
	fileBytes, err := utils.ReadFileData(file)
	if err != nil {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_UPLOAD)
	}

	// 创建文件记录
	fileModel := &model.File{
		FileName:    file.Filename,
		FileSize:    file.Size,
		ContentType: contentType,
		Data:        fileBytes,
		Hash:        hash,
		Status:      0,
		UploadTime:  time.Now().Unix(),
	}

	if err := s.repo.Create(fileModel); err != nil {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_UPLOAD)
	}
	return utils.JsonSuccessResp(ctx, fileModel)
}

func (s *UploadService) GetFile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	file, err := s.repo.FindByHash(id)
	if err != nil || file == nil {
		return utils.JsonErrorResp(ctx, model.ERROR_NOT_FOUND)
	}

	ctx.Set("Content-Type", file.ContentType)
	return ctx.Send(file.Data)
}

func (s *UploadService) DeleteFile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	file, err := s.repo.FindByHash(id)
	if err != nil || file == nil {
		return utils.JsonErrorResp(ctx, model.ERROR_NOT_FOUND)
	}

	if err := s.repo.Delete(file.ID); err != nil {
		return utils.JsonErrorResp(ctx, model.ERROR_FILE_DELETE)
	}

	return utils.JsonSuccessResp(ctx, nil)
}
