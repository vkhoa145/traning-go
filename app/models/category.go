package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserId      uint
}

type CreateCategoryInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Category) TableName() string {
	return "category"
}

func FilterCategoryRecord(cat *Category) *CategoryResponse {
	return &CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
	}
}
