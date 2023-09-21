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

func (Category) TableName() string {
	return "categories"
}

type CreateCategoryInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      uint   `json:"user_id" validates: "required"`
}

type UpdateCategoryInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      uint   `json:"user_id" validates: "required"`
}

type CategoryResponse struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
}

func FilterCategoryRecord(cat *Category) *CategoryResponse {
	return &CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
		UserId:      cat.UserId,
	}
}

var Categories []Category