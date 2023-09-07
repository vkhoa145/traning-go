package repository

import (
	"github.com/vkhoa145/go-training/app/models"
	"gorm.io/gorm"
)

type CategoryRepoInterface interface {
	CreateCategory(data *models.CreateCategoryInput, userId uint) (*models.Category, error)
}

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}
