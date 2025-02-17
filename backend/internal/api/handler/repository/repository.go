package repository

import (
	"LocalResourceHub/backend/internal/model"

	"gorm.io/gorm"
)

// GormFileRepository 实现 FileRepository 接口
type GormFileRepository struct {
	db *gorm.DB
}

// NewGormFileRepository 创建 GormFileRepository 实例
func NewGormFileRepository(db *gorm.DB) *GormFileRepository {
	db.AutoMigrate(&model.File{})
	return &GormFileRepository{db: db}
}

// Create 创建文件记录
func (r *GormFileRepository) Create(file *model.File) error {
	return r.db.Create(file).Error
}

// Delete 删除文件记录
func (r *GormFileRepository) Delete(id uint) error {
	return r.db.Delete(&model.File{}, id).Error
}

// FindByID 根据ID查找文件
func (r *GormFileRepository) FindByID(id uint) (*model.File, error) {
	var file model.File
	err := r.db.First(&file, id).Error
	return &file, err
}

// FindByHash 根据哈希值查找文件
func (r *GormFileRepository) FindByHash(hash string) (*model.File, error) {
	var file model.File
	err := r.db.Where("hash = ?", hash).First(&file).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &file, err
}

// List 获取文件列表
func (r *GormFileRepository) List(page, pageSize int, fileType string) ([]model.File, int64, error) {
	var files []model.File
	var total int64

	query := r.db.Model(&model.File{})

	if fileType != "all" {
		query = query.Where("content_type LIKE ?", fileType+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&files).Error

	return files, total, err
}
