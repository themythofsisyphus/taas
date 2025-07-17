package utils

import (
	"taas/model"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(context *gin.Context, statusCode int, message string, data interface{}) {
	response := model.APIResponse{
		Message: message,
		Success: true,
		Data:    data,
	}

	context.JSON(statusCode, response)
}

func ErrorResponse(context *gin.Context, statusCode int, message string, err interface{}) {
	respsonse := model.APIResponse{
		Message: message,
		Success: false,
		Error:   err,
	}

	context.JSON(statusCode, respsonse)
}
