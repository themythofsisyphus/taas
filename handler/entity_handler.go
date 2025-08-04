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
		return
	}

	newEntity, err := h.entityService.CreateEntity(context, &createEntityRequest)
	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Entity can't be created", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusCreated, "Entity Created Successfully", newEntity)
}

func (h *EntityHandler) DeleteEntity(context *gin.Context) {
	entityType := context.Param("type")

	if err := h.entityService.DeleteEntity(context, entityType); err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Entity can't be deleted", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusNoContent, "Entity is deleted", nil)
}

func (h *EntityHandler) ListEntities(context *gin.Context) {

	entities, err := h.entityService.GetAllEntities(context)

	if err != nil {
		utils.ErrorResponse(context, http.StatusInternalServerError, "Can't retrive entities", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tag retrived", entities)
}
