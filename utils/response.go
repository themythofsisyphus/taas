package utils

import (
	"taas/model"

	"github.com/gin-gonic/gin"
)

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

func ErrorResponse(context *gin.Context, statusCode int, message string, err any) {
	respsonse := model.APIResponse{
		Message: message,
		Success: false,
		Error:   err,
	}

	context.JSON(statusCode, respsonse)
}

func PaginationMetaResponse(total uint, limit int) *model.PaginationMetaResponse {
	return &model.PaginationMetaResponse{
		Total:      total,
		TotalPages: (int(total) + (limit - 1)) / limit,
	}
}
