package model

import (
	"errors"
)

// InMemoryFileRepository is an in-memory implementation of FileRepository
type InMemoryFileRepository struct {
	files map[string]*File
}

// List retrieves a paginated list of files from the repository
func (r *InMemoryFileRepository) List(page, pageSize int, fileType string) ([]File, int64, error) {
	var files []File
	total := int64(len(r.files))

	start := (page - 1) * pageSize
	end := min(start + pageSize, int(total))

	for _, file := range r.files {
		if fileType == "" || file.ContentType == fileType {
			files = append(files, *file)
		}
	}

	return files[start:end], total, nil
}

// FindByID retrieves a file by its ID from the repository
func (r *InMemoryFileRepository) FindByID(id uint) (*File, error) {
	for _, file := range r.files {
		if uint(file.ID) == id {
			return file, nil
		}
	}
	return nil, errors.New("file not found")
}

// NewFileRepository creates a new instance of FileRepository
func NewFileRepository() FileRepository {
	return &InMemoryFileRepository{
		files: make(map[string]*File),
	}
}

// Create adds a new file to the repository
func (r *InMemoryFileRepository) Create(file *File) error {
	if _, exists := r.files[file.Hash]; exists {
		return errors.New("file already exists")
	}
	r.files[file.Hash] = file
	return nil
}

// FindByHash retrieves a file by its hash from the repository
func (r *InMemoryFileRepository) FindByHash(hash string) (*File, error) {
	file, exists := r.files[hash]
	if !exists {
		return nil, errors.New("file not found")
	}
	return file, nil
}

// Delete removes a file from the repository by its ID
func (r *InMemoryFileRepository) Delete(id uint) error {
	for hash, file := range r.files {
		if uint(file.ID) == id {
			delete(r.files, hash)
			return nil
		}
	}
	return errors.New("file not found")
}
