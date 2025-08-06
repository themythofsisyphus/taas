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

// TagHandler handles HTTP requests related to tag operations.
type TagHandler struct {
	tagService *service.TagService
}

// NewTagHandler creates a new instance of TagHandler.
func NewTagHandler(service *service.TagService) *TagHandler {
	return &TagHandler{
		tagService: service,
	}
}

// ListTags returns a paginated list of tags.
func (h *TagHandler) ListTags(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))

	tags, err := h.tagService.GetTagsWithPagination(ctx, utils.NewPagination(page, limit))
	tagsCount, _ := h.tagService.GetTagsCount(ctx)

	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve tags", err.Error())
		return
	}

	meta := utils.PaginationMetaResponse(tagsCount, limit)
	utils.SuccessResponse(ctx, http.StatusOK, "Tags retrieved successfully", tags, meta)
}

// CreateTag handles the creation of a new tag.
func (h *TagHandler) CreateTag(ctx *gin.Context) {
	var req model.TagRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	newTag, err := h.tagService.CreateTag(ctx, &req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Failed to create tag", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Tag created successfully", newTag)
}

// UpdateTag updates an existing tag by ID.
func (h *TagHandler) UpdateTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid tag ID", err.Error())
		return
	}

	var req model.TagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	updatedTag, err := h.tagService.UpdateTag(ctx, uint(id), &req)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Failed to update tag", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Tag updated successfully", updatedTag)
}

// DeleteTag deletes a tag by ID.
func (h *TagHandler) DeleteTag(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid tag ID", err.Error())
		return
	}

	if err := h.tagService.DeleteTag(ctx, uint(id)); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Failed to delete tag", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Tag deleted successfully", nil)
}

// GetTagByID retrieves a single tag by ID.
func (h *TagHandler) GetTagByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid tag ID", err.Error())
		return
	}

	tag, err := h.tagService.GetTagByID(ctx, uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Tag not found", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Tag retrieved successfully", tag)
}

// SearchTags searches tags by a term with pagination.
func (h *TagHandler) SearchTags(ctx *gin.Context) {
	term := ctx.DefaultQuery("term", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))

	tags, err := h.tagService.SearchTags(ctx, term, utils.NewPagination(page, limit))
	tagsCount, _ := h.tagService.SearchTagsCount(ctx, term)

	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to search tags", err.Error())
		return
	}

	meta := utils.PaginationMetaResponse(tagsCount, limit)
	utils.SuccessResponse(ctx, http.StatusOK, "Tags retrieved successfully", tags, meta)
}
