package utils

import (
	"LocalResourceHub/backend/internal/model"
	"crypto/md5"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gofiber/fiber/v2"
)

func JsonSuccessResp(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(SuccessResponse(data))
}

func JsonErrorResp(ctx *fiber.Ctx, code int) error {
	// Always return HTTP status 200 for consistency
	return ctx.Status(fiber.StatusOK).JSON(ErrorResponse(code))
}

func SuccessResponse(data any) fiber.Map {
	resp := model.NewSuccessResponse(data)
	return fiber.Map{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Data,
	}
}

func ErrorResponse(code int) fiber.Map {
	resp := model.NewErrorResponse(code)
	return fiber.Map{
		"code":    resp.Code,
		"message": resp.Message,
		"data":    resp.Data,
	}
}

func ReadFileData(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return io.ReadAll(src)
}

// 获取环境变量或默认值
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 字符串反转
func ReverseString(s string) string {
	return strutil.Reverse(s)
}

// UUID生成
func GenerateUUID() (uuid string) {
	uuid, _ = random.UUIdV4()
	return uuid
}

// 计算文件哈希值
func CalculateFileHash(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", err
	}

	return cryptor.Sha256(string(hash.Sum(nil))), nil
}

// 生成分页数据
func Paginate(total, limit, page int) map[string]int {
	return map[string]int{
		"total":  total,
		"limit":  limit,
		"page":   page,
		"offset": (page - 1) * limit,
	}
}

// 获取环境变量
func GetEnv(key string, defaultValue string) string {
	return os.Getenv(key)
}

// 获取当前工作目录
func GetCurrentDir() (string, error) {
	return os.Getwd()
}

// 创建目录
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// 删除文件
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// 获取文件扩展名
func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}
