package models

import (
	"time"
)

// Service represents a service entity in the system
type Service struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Versions    int       `json:"versions" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ServiceResponse represents the response structure for service listings
type ServiceResponse struct {
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	Services []Service `json:"services"`
}

// SearchParams represents search and pagination parameters
type SearchParams struct {
	Page     int    `form:"page,default=1" binding:"min=1"`
	PageSize int    `form:"page_size,default=10" binding:"min=1,max=100"`
	Search   string `form:"search"`
	SortBy   string `form:"sort_by" binding:"omitempty,oneof=name -name"`
}
