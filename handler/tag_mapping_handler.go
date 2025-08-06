// Package handler defines HTTP handlers for managing entities, tags, tenants, and their mappings.
package handler

import (
	"net/http"
	"strconv"
	"taas/model"
	"taas/service"
	"taas/utils"

	"github.com/gin-gonic/gin"
)

// TagMappingHandler handles HTTP requests related to tag mapping operations.
type TagMappingHandler struct {
	tagMappingService *service.TagMappingService
}

// NewTagMappingHandler returns a new instance of TagMappingHandler.
func NewTagMappingHandler(service *service.TagMappingService) *TagMappingHandler {
	return &TagMappingHandler{
		tagMappingService: service,
	}
}

// CreateTagMappings handles the creation of tag mappings for a given entity.
func (h *TagMappingHandler) CreateTagMappings(ctx *gin.Context) {
	entityType := ctx.Param("entity_type")
	entityID, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid entity ID", err.Error())
		return
	}

	var tagMappingReq model.TagMappingRequest
	if err := ctx.ShouldBindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	tags, err := h.tagMappingService.CreateTagMappings(ctx, &tagMappingReq, entityType, uint(entityID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create tag mappings", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Tag mappings created successfully", tags)
}

// ListTagMappings handles listing tag mappings for a given entity with pagination.
func (h *TagMappingHandler) ListTagMappings(ctx *gin.Context) {
	entityType := ctx.Param("entity_type")
	entityID, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid entity ID", err.Error())
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	pagination := utils.NewPagination(page, limit)

	tags, err := h.tagMappingService.GetTagMappingsWithPagination(ctx, entityType, uint(entityID), pagination)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve tag mappings", err.Error())
		return
	}

	tagsCount, _ := h.tagMappingService.GetTagMappingsCount(ctx, entityType, uint(entityID))
	meta := utils.PaginationMetaResponse(tagsCount, limit)

	utils.SuccessResponse(ctx, http.StatusOK, "Tag mappings retrieved successfully", tags, meta)
}

// DeleteTagMappings handles the removal of tag mappings for a given entity.
func (h *TagMappingHandler) DeleteTagMappings(ctx *gin.Context) {
	entityType := ctx.Param("entity_type")
	entityID, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid entity ID", err.Error())
		return
	}

	var tagMappingReq model.TagMappingRequest
	if err := ctx.ShouldBindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err = h.tagMappingService.RemoveTagMappings(ctx, &tagMappingReq, entityType, uint(entityID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete tag mappings", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Tag mappings deleted successfully", gin.H{})
}
