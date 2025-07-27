package service

import (
	"taas/model"
	"taas/repository"
)

type TenantService struct {
	tenantRepository *repository.TenantRepository
}

func NewTenantService(repo *repository.TenantRepository) *TenantService {
	return &TenantService{tenantRepository: repo}
}

func (s *TenantService) CreateTenant(tenantRequest *model.Tenant) (*model.Tenant, error) {
	return s.tenantRepository.Create(tenantRequest)
}

func (s *TenantService) DeleteTenant(tenantId uint) error {
	return s.tenantRepository.Delete(tenantId)
}

func (s *TenantService) GetTenantByID(tenantID uint) (*model.Tenant, error) {
	return s.tenantRepository.GetByID(tenantID)
}
