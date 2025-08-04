package handler

import (
	"net/http"
	"taas/model"
	"taas/service"
	"taas/utils"

	"strconv"

	"github.com/gin-gonic/gin"
)

type TagMappingHandler struct {
	tagMappingService *service.TagMappingService
}

func NewTagMappingHandler(service *service.TagMappingService) *TagMappingHandler {
	return &TagMappingHandler{
		tagMappingService: service,
	}
}

func (h *TagMappingHandler) CreateTagMappings(context *gin.Context) {
	entityType := context.Param("entity_type")
	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param", err.Error())
		return
	}
	var tagMappingReq model.TagMappingRequest

	if err := context.BindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	tags, err := h.tagMappingService.CreateTagMappings(context, &tagMappingReq, entityType, uint(entityID))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Mappings retrived Successfully", tags)
}

func (h *TagMappingHandler) ListTagMappings(context *gin.Context) {
	entityType := context.Param("entity_type")
	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "50"))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param", err.Error())
		return
	}

	tags, err := h.tagMappingService.GetTagMappingsWithPagination(context, entityType, uint(entityID), utils.NewPagination(page, limit))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive", err.Error())
		return
	}

	tagsCount, _ := h.tagMappingService.GetTagMappingsCount(context, entityType, uint(entityID))

	metaResponse := utils.PaginationMetaResponse(tagsCount, limit)

	utils.SuccessResponse(context, http.StatusOK, "Tag Mappings retrived Successfully", tags, metaResponse)
}

func (h *TagMappingHandler) DeleteTagMappings(context *gin.Context) {
	entityType := context.Param("entity_type")
	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param", err.Error())
		return
	}
	var tagMappingReq model.TagMappingRequest

	if err := context.BindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err = h.tagMappingService.RemoveTagMappings(context, &tagMappingReq, entityType, uint(entityID))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Tag can't be deleted", nil)
		return
	}

	utils.SuccessResponse(context, http.StatusNoContent, "Tag mappings deleted", gin.H{})
}
