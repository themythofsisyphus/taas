// Package model contains shared data structures for API responses and pagination metadata.
package model

// APIResponse defines the standard structure of API responses.
// It supports embedding data, error messages, and metadata.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// PaginationMetaResponse provides metadata about paginated results,
// such as total count and number of pages.
type PaginationMetaResponse struct {
	Total      uint `json:"total"`
	TotalPages int  `json:"total_pages"`
}
