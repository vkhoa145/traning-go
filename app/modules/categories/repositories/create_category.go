package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) CreateCategory(data *models.CreateCategoryInput) (*models.Category, error) {
	category := &models.Category{
		Name:        data.Name,
		Description: data.Description,
		UserId:      1,
	}

	result := r.DB.Table("category").Create(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
