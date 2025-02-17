package model

import (
	"mime/multipart"

	"gorm.io/gorm"
)

// 支持的文件类型
var AllowedFileTypes = []string{
	"application/pdf",
	"application/msword", // .doc
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // .docx
	"application/vnd.ms-excel", // .xls
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", // .xlsx
	"text/plain",
	"text/csv",
	"application/zip",
	"application/x-rar-compressed",
}

type (
	// FileValidator interface for validating file properties
	FileValidator interface {
		ValidateFileSize(size int64) error
		ValidateFileType(contentType string, fileType string) error
	}

	// FileStorage interface for handling file storage operations
	FileStorage interface {
		processFile(file *multipart.FileHeader) (*File, error)
		// SaveFile(file *multipart.FileHeader, uuid string) (string, error)
		// GenerateFilePath(uuid string, ext string) string
	}

	// FileRepository interface for interacting with the file data store
	FileRepository interface {
		Create(file *File) error
		Delete(id uint) error
		FindByID(id uint) (*File, error)
		FindByHash(hash string) (*File, error)
		List(page, pageSize int, fileType string) ([]File, int64, error)
	}
)

// 文件基础结构体
type (
	File struct {
		gorm.Model
		FileName    string `json:"file_name"`    // 文件名
		FileSize    int64  `json:"file_size"`    // 文件大小(字节)
		ContentType string `json:"content_type"` // 文件类型
		Data        []byte `json:"-"`            // 文件二进制数据
		Hash        string `json:"hash"`         // 哈希值
		Status      int    `json:"status"`       // 状态(0:正常,1:已删除)
		UploadTime  int64  `json:"upload_time"`  // 上传时间戳
	}
	FileList struct {
		List []File `json:"list"` // 文件列表
	}
)
