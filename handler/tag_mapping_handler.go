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
	}
	var tagMappingReq model.TagMappingRequest

	if err := context.BindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid request", err.Error())
	}

	tags, err := h.tagMappingService.CreateTagMappings(context, &tagMappingReq, entityType, uint(entityID))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive", nil)
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Mappings retrived Successfully", tags)
}

func (h *TagMappingHandler) ListTagMappings(context *gin.Context) {
	entityType := context.Param("entity_type")
	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param", err.Error())
	}

	tags, err := h.tagMappingService.GetTagMappings(context, entityType, uint(entityID))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive", nil)
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Mappings retrived Successfully", tags)
}

func (h *TagMappingHandler) DeleteTagMappings(context *gin.Context) {
	entityType := context.Param("entity_type")
	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param", err.Error())
	}
	var tagMappingReq model.TagMappingRequest

	if err := context.BindJSON(&tagMappingReq); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid request", err.Error())
	}

	err = h.tagMappingService.RemoveTagMappings(context, &tagMappingReq, entityType, uint(entityID))

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Tag can't be deleted", nil)
	}

	utils.SuccessResponse(context, http.StatusNoContent, "Tag mappings deleted", gin.H{})
}

// func RegisterTagMappingRoutes(route *gin.Engine, db *gorm.DB) {

// 	repo := repository.NewTagMappingRepo(db)
// 	entityService := service.NewEntityService(repository.NewEntityRepo(db))
// 	tagService := service.NewTagService(repository.NewTagRepository(db))
// 	service := service.NewTagMappingService(repo, entityService, tagService)

// 	mappingGroup := route.Group(":entity_type/tag_mappings")

// 	mappingGroup.POST("/:id", func(c *gin.Context) { createTagMappings(c, service) })
// 	mappingGroup.DELETE("/:id", func(c *gin.Context) { deleteTagMappings(c, service) })
// 	mappingGroup.GET("/:id", func(c *gin.Context) { getTagMappings(c, service) })
// }

// // private methods

// func createTagMappings(context *gin.Context, service *service.TagMappingService) {
// 	entityType := context.Param("entity_type")
// 	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)
// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid entity ID"})
// 		return
// 	}

// 	var tagIDs []uint
// 	if err := context.ShouldBindJSON(&tagIDs); err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	service.CreateTagMappings(context, tagIDs, entityType, uint(entityID))
// 	context.JSON(201, nil)
// }

// func deleteTagMappings(context *gin.Context, service *service.TagMappingService) {
// 	entityType := context.Param("entity_type")
// 	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid entity ID"})
// 		return
// 	}

// 	var tagIDs []uint
// 	if err := context.ShouldBindJSON(&tagIDs); err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	service.RemoveTagMappings(context, tagIDs, entityType, uint(entityID))
// 	context.JSON(204, gin.H{})
// }

// func getTagMappings(context *gin.Context, service *service.TagMappingService) {
// 	entityType := context.Param("entity_type")
// 	entityID, err := strconv.ParseUint(context.Param("id"), 10, 0)

// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid entity ID"})
// 		return
// 	}

// 	tags, err := service.GetTagMappings(context, entityType, uint(entityID))

// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid entity ID"})
// 		return
// 	}
// 	context.JSON(200, tags)
// }
