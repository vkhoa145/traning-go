package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) CreateCategory(data *models.CreateCategoryInput, userId uint) (*models.Category, error) {
	category := &models.Category{
		Name:        data.Name,
		Description: data.Description,
		UserId:      userId,
	}

	result := r.DB.Table(models.Category{}.TableName()).Create(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
