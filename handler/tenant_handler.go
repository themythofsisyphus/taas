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

// TenantHandler handles tenant-related API operations.
type TenantHandler struct {
	tenantService *service.TenantService
}

// NewTenantHandler creates a new instance of TenantHandler.
func NewTenantHandler(service *service.TenantService) *TenantHandler {
	return &TenantHandler{tenantService: service}
}

// CreateTenant handles the creation of a new tenant.
func (h *TenantHandler) CreateTenant(ctx *gin.Context) {
	var tenantReq model.TenantRecord

	if err := ctx.ShouldBindJSON(&tenantReq); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	newTenant := model.Tenant{
		ID: tenantReq.TenantID,
	}

	createdTenant, err := h.tenantService.CreateTenant(&newTenant)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create tenant", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Tenant created successfully", createdTenant)
}

// GetTenantByID handles fetching a tenant by ID.
func (h *TenantHandler) GetTenantByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid tenant ID", err.Error())
		return
	}

	tenant, err := h.tenantService.GetTenantByID(uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Tenant not found", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Tenant retrieved successfully", tenant)
}

// DeleteTenant handles deletion of a tenant.
func (h *TenantHandler) DeleteTenant(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid tenant ID", err.Error())
		return
	}

	if err := h.tenantService.DeleteTenant(uint(id)); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete tenant", err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Tenant deleted successfully", nil)
}
