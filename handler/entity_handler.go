// Package handler defines HTTP handlers for managing entities, tags, tenants, and their mappings.
package handler

import (
	"net/http"
	"taas/model"
	"taas/service"
	"taas/utils"

	"github.com/gin-gonic/gin"
)

// EntityHandler handles HTTP requests related to entity operations.
type EntityHandler struct {
	entityService *service.EntityService
}

// NewEntityHandler creates a new instance of EntityHandler.
func NewEntityHandler(service *service.EntityService) *EntityHandler {
	return &EntityHandler{
		entityService: service,
	}
}

// CreateEntity handles the creation of a new entity.
func (h *EntityHandler) CreateEntity(ctx *gin.Context) {
	var req model.EntityRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	newEntity, err := h.entityService.CreateEntity(ctx, &req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create entity", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Entity created successfully", newEntity)
}

// DeleteEntity handles deletion of an entity by its type.
func (h *EntityHandler) DeleteEntity(ctx *gin.Context) {
	entityType := ctx.Param("type")

	if err := h.entityService.DeleteEntity(ctx, entityType); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete entity", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Entity deleted successfully", nil)
}

// ListEntities returns all registered entities.
func (h *EntityHandler) ListEntities(ctx *gin.Context) {
	entities, err := h.entityService.GetAllEntities(ctx)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve entities", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Entities retrieved successfully", entities)
}
