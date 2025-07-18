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

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrive tags", err.Error())
	}

	utils.SuccessResponse(context, http.StatusOK, "Tags retrived successfully", tags)
}

func (h *TagHandler) CreateTag(context *gin.Context) {
	var createTagRequest model.TagRequest

	if err := context.BindJSON(&createTagRequest); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	newTag, err := h.tagService.CreateTag(context, &createTagRequest)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	utils.SuccessResponse(context, http.StatusCreated, "Tag Created Successfully", newTag)
}

func (h *TagHandler) UpdateTag(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
	}
	var updateTagRequest model.TagRequest

	if err := context.BindJSON(&updateTagRequest); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	updatedTag, err := h.tagService.UpdateTag(context, uint(id), &updateTagRequest)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Updated Successfully", updatedTag)
}

func (h *TagHandler) DeleteTag(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
	}

	err = h.tagService.DeleteTag(context, uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Tag can't be deleted", err.Error())
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Deleted", nil)
}

func (h *TagHandler) GetTagByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Param ID", err.Error())
	}

	tag, err := h.tagService.GetTagByID(context, uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Tag can't be deleted", err.Error())
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag Retrived", tag)
}

// func RegisterTagRoutes(route *gin.Engine, db *gorm.DB) {

// 	repo := repository.NewTagRepository(db)
// 	service := service.NewTagService(repo)

// 	tagGroup := route.Group("/tags")
// 	tagGroup.GET("", func(c *gin.Context) { getTags(c, service) })
// 	tagGroup.POST("", func(c *gin.Context) { createTag(c, service) })
// 	tagGroup.GET("/:id", func(c *gin.Context) { getTagByID(c, service) })
// 	tagGroup.PUT("/:id", func(c *gin.Context) { updateTag(c, service) })
// 	tagGroup.DELETE("/:id", func(c *gin.Context) { deleteTag(c, service) })
// }

// // private methods

// func getTags(context *gin.Context, service *service.TagService) {
// 	tags, err := service.GetAllTags(context)
// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to fetch tags"})
// 		return
// 	}
// 	context.JSON(200, tags)
// }

// func createTag(context *gin.Context, service *service.TagService) {
// 	var tag model.TagRequest
// 	if err := context.ShouldBindJSON(&tag); err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	tagModel := model.Tag{
// 		Name: tag.Name,
// 	}

// 	createdTag, err := service.CreateTag(context, &tagModel)
// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to create tag"})
// 		return
// 	}
// 	context.JSON(201, createdTag)

// }
// func getTagByID(context *gin.Context, service *service.TagService) {
// 	id := context.Param("id")
// 	tagId, err := strconv.ParseUint(id, 10, 0)
// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid tag ID"})
// 		return
// 	}
// 	tag, err := service.GetTagByID(context, uint(tagId))
// 	if err != nil {
// 		context.JSON(404, gin.H{"error": "Tag not found"})
// 		return
// 	}
// 	context.JSON(200, tag)
// }
// func updateTag(context *gin.Context, service *service.TagService) {
// 	id := context.Param("id")
// 	tagId, err := strconv.ParseUint(id, 10, 0)
// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid tag ID"})
// 		return
// 	}

// 	var tag dto.TagRequest
// 	if err := context.ShouldBindJSON(&tag); err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	existingTag, err := service.GetTagByID(context, uint(tagId))
// 	if err != nil {
// 		context.JSON(404, gin.H{"error": "Tag not found"})
// 		return
// 	}

// 	existingTag.Name = tag.Name

// 	updatedTag, err := service.UpdateTag(context, &existingTag)
// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to update tag"})
// 		return
// 	}
// 	context.JSON(200, updatedTag)

// }
// func deleteTag(context *gin.Context, service *service.TagService) {
// 	id := context.Param("id")
// 	tagId, err := strconv.ParseUint(id, 10, 0)
// 	if err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid tag ID"})
// 		return
// 	}
// 	err = service.DeleteTag(context, uint(tagId))
// 	if err != nil {
// 		context.JSON(404, gin.H{"error": "Tag not found"})
// 		return
// 	}
// 	context.JSON(204, gin.H{})
// }
