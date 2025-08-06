// Package utils contains utility functions including caching logic and helpers.
package utils

import (
	"taas/model"

	"github.com/gin-gonic/gin"
)

// SuccessResponse sends a standard JSON response with success = true.
// Accepts optional meta data as a variadic parameter.
func SuccessResponse(context *gin.Context, statusCode int, message string, data any, meta ...interface{}) {
	var metaValue any
	if len(meta) > 0 {
		metaValue = meta[0]
	}

	response := model.APIResponse{
		Message: message,
		Success: true,
		Data:    data,
		Meta:    metaValue,
	}

	context.JSON(statusCode, response)
}

// ErrorResponse sends a standard JSON response with success = false.
func ErrorResponse(context *gin.Context, statusCode int, message string, err any) {
	response := model.APIResponse{
		Message: message,
		Success: false,
		Error:   err,
	}

	context.JSON(statusCode, response)
}

// PaginationMetaResponse builds pagination metadata for paginated endpoints.
func PaginationMetaResponse(total uint, limit int) *model.PaginationMetaResponse {
	return &model.PaginationMetaResponse{
		Total:      total,
		TotalPages: (int(total) + (limit - 1)) / limit,
	}
}
