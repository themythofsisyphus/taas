package handler

import (
	"net/http"
	"strconv"
	"taas/model"
	"taas/service"
	"taas/utils"

	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	tenantService *service.TenantService
}

func NewTenantHandler(service *service.TenantService) *TenantHandler {
	return &TenantHandler{tenantService: service}
}

func (h *TenantHandler) CreateTenant(context *gin.Context) {
	var tenant model.TenantRecord

	if err := context.BindJSON(&tenant); err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	newTenant := model.Tenant{
		ID: tenant.TenantID,
	}

	createdTenant, err := h.tenantService.CreateTenant(&newTenant)

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
	}

	utils.SuccessResponse(context, http.StatusCreated, "Tenant Created Successfully", createdTenant)
}

func (h *TenantHandler) GetTenantByID(context *gin.Context) {
	tenantID := context.Param("id")

	id, err := strconv.ParseUint(tenantID, 10, 32)
	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Tenant ID", err.Error())
		return
	}

	tenant, err := h.tenantService.GetTenantByID(uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tenant retrieved successfully", tenant)
}

func (h *TenantHandler) DeleteTenant(context *gin.Context) {
	tenantID := context.Param("id")

	id, err := strconv.ParseUint(tenantID, 10, 32)
	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Tenant ID", err.Error())
		return
	}

	err = h.tenantService.DeleteTenant(uint(id))

	if err != nil {
		utils.ErrorResponse(context, http.StatusBadRequest, "Invalid Request", err.Error())
		return
	}

	utils.SuccessResponse(context, http.StatusOK, "Tenant deleted successfully", nil)
}
