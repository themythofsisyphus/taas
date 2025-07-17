package handler

import (
	"net/http"
	"taas/model"
	"taas/service"
	"taas/utils"

	"github.com/gin-gonic/gin"
)

type EntityHandler struct {
	entityService *service.EntityService
}

func NewEntityHandler(service *service.EntityService) *EntityHandler {
	return &EntityHandler{
		entityService: service,
	}
}

func (h *EntityHandler) CreateEntity(context *gin.Context) {
	var createEntityRequest model.EntityRequest

	if err := context.BindJSON(&createEntityRequest); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	newEntity, err := h.entityService.CreateEntity(context, &createEntityRequest)
	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Entity can't be created", err.Error())
	}

	utils.SuccessResponse(context, http.StatusCreated, "Entity Created Successfully", newEntity)
}

func (h *EntityHandler) DeleteEntity(context *gin.Context) {
	entityType := context.Param("type")

	if err := h.entityService.DeleteEntity(context, entityType); err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Entity can't be deleted", err.Error())
	}

	utils.SuccessResponse(context, http.StatusNoContent, "Entity is deleted", nil)
}

func (h *EntityHandler) ListEntities(context *gin.Context) {

	entities, err := h.entityService.GetAllEntities(context)

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive entities", err.Error())
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag retrived", entities)
}

// func RegisterEntityRoutes(route *gin.Engine, db *gorm.DB) {

// 	repo := repository.NewEntityRepo(db)
// 	service := service.NewEntityService(repo)

// 	entityRoutes := route.Group("/entities")
// 	entityRoutes.GET("", func(c *gin.Context) { getEntities(c, service) })
// 	entityRoutes.POST("", func(c *gin.Context) { createEntity(c, service) })
// 	entityRoutes.DELETE("/:type", func(c *gin.Context) { deleteEntity(c, service) })
// }

// func getEntities(context *gin.Context, service *service.EntityService) {
// 	ctx := context.Request.Context()
// 	entities, err := service.GetAllEntities(ctx)
// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to fetch entities"})
// 		return
// 	}

// 	context.JSON(200, entities)
// }

// func createEntity(context *gin.Context, service *service.EntityService) {
// 	var entity model.Entity
// 	ctx := context.Request.Context()

// 	if err := context.ShouldBindJSON(&entity); err != nil {
// 		context.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	newEntity, err := service.CreateEntity(ctx, &entity)

// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to create entity"})
// 		return
// 	}

// 	context.JSON(201, newEntity)
// }

// func deleteEntity(context *gin.Context, service *service.EntityService) {
// 	entityType := context.Param("type")
// 	ctx := context.Request.Context()

// 	err := service.DeleteEntity(ctx, entityType)

// 	if err != nil {
// 		context.JSON(500, gin.H{"error": "Failed to delete entity"})
// 		return
// 	}

// 	context.JSON(204, gin.H{})
// }
