package handler

import (
	"net/http"
	"taas/service"
	"taas/utils"

	"strconv"
	"taas/model"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler(service *service.TagService) *TagHandler {
	return &TagHandler{
		tagService: service,
	}
}

func (h *TagHandler) ListTags(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "50"))

	tags, err := h.tagService.GetTagsWithPagination(context, utils.NewPagination(page, limit))
	tagsCount, _ := h.tagService.GetTagsCount(context)

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrive tags", err.Error())
		return
	}

	metaResponse := utils.PaginationMetaResponse(tagsCount, limit)

	utils.SuccessResponse(context, http.StatusOK, "Tags retrived successfully", tags, metaResponse)
}

func (h *TagHandler) CreateTag(context *gin.Context) {
	var createTagRequest model.TagRequest

	if err := context.BindJSON(&createTagRequest); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	newTag, err := h.tagService.CreateTag(context, &createTagRequest)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusCreated, "Tag Created Successfully", newTag)
}

func (h *TagHandler) UpdateTag(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
		return
	}
	var updateTagRequest model.TagRequest

	if err := context.BindJSON(&updateTagRequest); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	updatedTag, err := h.tagService.UpdateTag(context, uint(id), &updateTagRequest)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Updated Successfully", updatedTag)
}

func (h *TagHandler) DeleteTag(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
		return
	}

	err = h.tagService.DeleteTag(context, uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Tag can't be deleted", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Deleted", nil)
}

func (h *TagHandler) GetTagByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
		return
	}

	tag, err := h.tagService.GetTagByID(context, uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusNotFound, "Tag not found", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Retrived", tag)
}

func (h *TagHandler) SearchTags(context *gin.Context) {
	searchTerm := context.DefaultQuery("term", "")
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "50"))

	tags, err := h.tagService.SearchTags(context, searchTerm, utils.NewPagination(page, limit))
	tagsCount, _ := h.tagService.SearchTagsCount(context, searchTerm)

	if err != nil {
		utils.ErrorResponse(context, http.StatusNotFound, "Tag not found", err.Error())
		return
	}

	metaResponse := utils.PaginationMetaResponse(tagsCount, limit)

	utils.SuccessResponse(context, http.StatusOK, "Tag Retrived", tags, metaResponse)
}
