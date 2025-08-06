package service

import (
	"taas/model"
	"taas/repository"
)

// TenantService handles business logic for tenants.
type TenantService struct {
	tenantRepository *repository.TenantRepository
}

// NewTenantService creates a new instance of TenantService.
func NewTenantService(repo *repository.TenantRepository) *TenantService {
	return &TenantService{tenantRepository: repo}
}

// CreateTenant creates a new tenant.
func (s *TenantService) CreateTenant(tenant *model.Tenant) (*model.Tenant, error) {
	return s.tenantRepository.Create(tenant)
}

// DeleteTenant removes a tenant by ID.
func (s *TenantService) DeleteTenant(tenantID uint) error {
	return s.tenantRepository.Delete(tenantID)
}

// GetTenantByID retrieves a tenant by ID.
func (s *TenantService) GetTenantByID(tenantID uint) (*model.Tenant, error) {
	return s.tenantRepository.GetByID(tenantID)
}
