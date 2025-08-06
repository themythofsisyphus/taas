// Package repository provides access to tenant-related database operations.
package repository

import (
	"taas/model"

	"gorm.io/gorm"
)

// TenantRepository handles database operations related to tenants.
type TenantRepository struct {
	db *gorm.DB
}

// NewTenantRepository returns a new instance of TenantRepository.
func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

// Create adds a new tenant to the database.
func (r *TenantRepository) Create(tenant *model.Tenant) (*model.Tenant, error) {
	err := r.db.Create(tenant).Error
	return tenant, err
}

// Delete removes a tenant by ID.
func (r *TenantRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tenant{}, id).Error
}

// GetByID retrieves a tenant by ID.
func (r *TenantRepository) GetByID(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.First(&tenant, id).Error
	return &tenant, err
}
