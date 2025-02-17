package model

// General response structure
type Response struct {
	Code    int    `json:"code"`    // Status code
	Message string `json:"message"` // Prompt information
	Data    any    `json:"data"`    // Data
}

// File upload response
type FileUploadResponse struct {
	Response
	FileInfo any `json:"file_info"` // File information
}

// Status code definitions
const (
	// Success status code
	SUCCESS = 200

	// Client error status codes (4xx)
	ERROR_INVALID_REQUEST  = 400 // Invalid request
	ERROR_UNAUTHORIZED     = 401 // Unauthorized
	ERROR_FORBIDDEN        = 403 // Forbidden
	ERROR_NOT_FOUND        = 404 // Resource not found
	ERROR_FILE_TOO_LARGE   = 413 // File too large
	ERROR_UNSUPPORTED_TYPE = 415 // Unsupported file type

	// Server error status codes (5xx)
	ERROR_INTERNAL        = 500 // Internal server error
	ERROR_NOT_IMPLEMENTED = 501 // Not implemented
	ERROR_STORAGE_FULL    = 507 // Insufficient storage
)

// File upload related error codes
const (
	ERROR_FILE_EMPTY        = 4001 // Empty file
	ERROR_FILE_UPLOAD       = 4002 // File upload failed
	ERROR_FILE_DELETE       = 4003 // File deletion failed
	ERROR_FILE_DOWNLOAD     = 4004 // File download failed
	ERROR_FILE_CORRUPTED    = 4005 // File corrupted
	ERROR_FILE_EXISTS       = 4006 // File already exists
	ERROR_FILE_NAME_INVALID = 4007 // Invalid file name
	ERROR_FILE_NOT_FOUND    = 4008 // File not found
)

// Image related error codes
const (
	ERROR_IMAGE_FORMAT     = 4101 // Image format error
	ERROR_IMAGE_SIZE       = 4102 // Image size error
	ERROR_IMAGE_CORRUPTED  = 4103 // Image corrupted
	ERROR_IMAGE_PROCESSING = 4104 // Image processing failed
)

// Error message mapping
var ErrorMessages = map[int]string{
	SUCCESS:                "Operation successful",
	ERROR_INVALID_REQUEST:  "Invalid request parameters",
	ERROR_UNAUTHORIZED:     "Please log in first",
	ERROR_FORBIDDEN:        "Permission denied",
	ERROR_NOT_FOUND:        "Requested resource not found",
	ERROR_FILE_TOO_LARGE:   "File size exceeds limit",
	ERROR_UNSUPPORTED_TYPE: "Unsupported file type",
	ERROR_INTERNAL:         "Internal server error",
	ERROR_NOT_IMPLEMENTED:  "Interface not implemented",
	ERROR_STORAGE_FULL:     "Insufficient storage",
	ERROR_FILE_EMPTY:       "Uploaded file is empty",
	ERROR_FILE_UPLOAD:      "File upload failed",
	ERROR_FILE_DELETE:      "File deletion failed",
	ERROR_FILE_DOWNLOAD:    "File download failed",
	ERROR_FILE_CORRUPTED:   "File is corrupted",
	ERROR_FILE_EXISTS:      "File already exists",
	ERROR_FILE_NAME_INVALID: "File name contains invalid characters",
	ERROR_IMAGE_FORMAT:     "Unsupported image format",
	ERROR_IMAGE_SIZE:       "Image dimensions do not meet requirements",
	ERROR_IMAGE_CORRUPTED:  "Image file is corrupted",
	ERROR_IMAGE_PROCESSING: "Image processing failed",
	ERROR_FILE_NOT_FOUND:   "File not found",
}

// Create success response
func NewSuccessResponse(data any) Response {
	return Response{
		Code:    SUCCESS,
		Message: ErrorMessages[SUCCESS],
		Data:    data,
	}
}

// Create error response
func NewErrorResponse(code int) Response {
	message, exists := ErrorMessages[code]
	if !exists {
		message = "Unknown error"
	}
	return Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// Create file upload success response
func NewFileUploadSuccessResponse(fileInfo any) FileUploadResponse {
	return FileUploadResponse{
		Response: NewSuccessResponse(nil),
		FileInfo: fileInfo,
	}
}
