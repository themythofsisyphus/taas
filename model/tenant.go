// Package model contains data structures for tenant identification and metadata.
package model

// Tenant represents a tenant entity with a unique ID.
type Tenant struct {
	ID uint `gorm:"primaryKey;autoIncrement:false"`
}

// TenantRecord defines a lightweight representation of a tenant for API responses or request payloads.
type TenantRecord struct {
	TenantID uint `json:"tenant_id"`
}
