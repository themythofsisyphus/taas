package handlers

import (
	"taas/dto"
	"taas/repository"
	"taas/services"

	"strconv"
	"taas/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTagRoutes(route *gin.Engine, db *gorm.DB) {

	repo := repository.NewTagRepository(db)
	service := services.NewTagService(repo)

	tagGroup := route.Group("/tags")
	tagGroup.GET("", func(c *gin.Context) { getTags(c, service) })
	tagGroup.POST("", func(c *gin.Context) { createTag(c, service) })
	tagGroup.GET("/:id", func(c *gin.Context) { getTagByID(c, service) })
	tagGroup.PUT("/:id", func(c *gin.Context) { updateTag(c, service) })
	tagGroup.DELETE("/:id", func(c *gin.Context) { deleteTag(c, service) })
}

// private methods

func getTags(context *gin.Context, service services.TagService) {
	tags, err := service.GetAllTags(currentTenantID(context))
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to fetch tags"})
		return
	}
	context.JSON(200, tags)
}

func createTag(context *gin.Context, service services.TagService) {
	var tag dto.TagRequest
	tag.TenantID = currentTenantID(context)
	if err := context.ShouldBindJSON(&tag); err != nil {
		context.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	tagModel := models.Tag{
		Name:     tag.Name,
		TenantID: tag.TenantID,
	}

	createdTag, err := service.CreateTag(&tagModel)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to create tag"})
		return
	}
	context.JSON(201, createdTag)

}
func getTagByID(context *gin.Context, service services.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}
	tag, err := service.GetTagByID(uint(tagId), currentTenantID(context))
	if err != nil {
		context.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	context.JSON(200, tag)
}
func updateTag(context *gin.Context, service services.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}

	var tag dto.TagRequest
	tag.TenantID = currentTenantID(context)
	if err := context.ShouldBindJSON(&tag); err != nil {
		context.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	tagModel := models.Tag{
		ID:       uint(tagId),
		Name:     tag.Name,
		TenantID: tag.TenantID,
	}

	updatedTag, err := service.UpdateTag(&tagModel)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to update tag"})
		return
	}
	context.JSON(200, updatedTag)

}
func deleteTag(context *gin.Context, service services.TagService) {
	id := context.Param("id")
	tagId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid tag ID"})
		return
	}
	err = service.DeleteTag(uint(tagId), currentTenantID((context)))
	if err != nil {
		context.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	context.JSON(204, gin.H{})
}

func currentTenantID(context *gin.Context) uint {
	tenantID, exists := context.Get("tenant_id")
	if !exists {
		context.JSON(400, gin.H{"error": "Tenant ID not found"})
		return 0
	}
	return tenantID.(uint)
}
