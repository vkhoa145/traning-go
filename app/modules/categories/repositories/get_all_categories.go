package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) GetAllCategories(userId float64) ([]models.Category, error) {

	categories := models.Categories
	result := r.DB.Table(models.Category{}.TableName()).Where("user_id=?", userId).Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
