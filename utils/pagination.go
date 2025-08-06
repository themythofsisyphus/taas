// Package utils contains utility functions including caching logic and helpers.
package utils

import "gorm.io/gorm"

// Pagination represents pagination parameters for queries.
type Pagination struct {
	Page  int // Current page number
	Limit int // Number of records per page
}

// NewPagination validates and returns a Pagination object.
// If page < 1, it defaults to 1.
// If limit < 1 or > 50, it defaults to 10.
func NewPagination(page int, limit int) *Pagination {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 50 {
		limit = 10
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

// DBScope returns a GORM scope function to apply pagination to a query.
func (p *Pagination) DBScope(_ *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit)
	}
}
