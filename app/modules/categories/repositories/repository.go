package repository

import (
	"github.com/vkhoa145/go-training/app/models"
	"gorm.io/gorm"
)

type CategoryRepoInterface interface {
	CreateCategory(data *models.CreateCategoryInput) (*models.Category, error)
	GetCategoryById(id float64, userId float64) (*models.Category, error)
	UpdateCategory(data *models.UpdateCategoryInput, category *models.Category) (*models.Category, error)
	DeleteCategory(existedCategory *models.Category) (*models.Category, error)
	GetAllCategories(userId float64) ([]models.Category, error)
}

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}
