package handler

import (
	"taas/dto"
	"taas/repository"
	"taas/service"

	"strconv"
	"taas/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTagRoutes(route *gin.Engine, db *gorm.DB) {

	repo := repository.NewTagRepository(db)
	service := service.NewTagService(repo)

	tagGroup := route.Group("/tags")
	tagGroup.GET("", func(c *gin.Context) { getTags(c, service) })
	tagGroup.POST("", func(c *gin.Context) { createTag(c, service) })
	tagGroup.GET("/:id", func(c *gin.Context) { getTagByID(c, service) })
	tagGroup.PUT("/:id", func(c *gin.Context) { updateTag(c, service) })
	tagGroup.DELETE("/:id", func(c *gin.Context) { deleteTag(c, service) })
}

// private methods

func getTags(context *gin.Context, service *service.TagService) {
	tags, err := service.GetAllTags(context)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to fetch tags"})
		return
	}
	context.JSON(200, tags)
}

func createTag(context *gin.Context, service *service.TagService) {
	var tag dto.TagRequest
	if err := context.ShouldBindJSON(&tag); err != nil {
		context.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	tagModel := model.Tag{
		Name: tag.Name,
	}

	createdTag, err := service.CreateTag(context, &tagModel)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to create tag"})
		return
	}
	context.JSON(201, createdTag)

}
func getTagByID(context *gin.Context, service *service.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}
	tag, err := service.GetTagByID(context, uint(tagId))
	if err != nil {
		context.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	context.JSON(200, tag)
}
func updateTag(context *gin.Context, service *service.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}

	var tag dto.TagRequest
	if err := context.ShouldBindJSON(&tag); err != nil {
		context.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	existingTag, err := service.GetTagByID(context, uint(tagId))
	if err != nil {
		context.JSON(404, gin.H{"error": "Tag not found"})
		return
	}

	existingTag.Name = tag.Name

	updatedTag, err := service.UpdateTag(context, &existingTag)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to update tag"})
		return
	}
	context.JSON(200, updatedTag)

}
func deleteTag(context *gin.Context, service *service.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}
	err = service.DeleteTag(context, uint(tagId))
	if err != nil {
		context.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	context.JSON(204, gin.H{})
}
