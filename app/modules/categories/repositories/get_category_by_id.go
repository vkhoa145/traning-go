package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) GetCategoryById(id float64, userId float64) (*models.Category, error) {

	var category *models.Category
	result := r.DB.Table(models.Category{}.TableName()).Where("user_id=? AND id=?", userId, id).First(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
