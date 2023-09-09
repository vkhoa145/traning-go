package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) DeleteCategory(existedCategory *models.Category) (*models.Category, error) {

	result := r.DB.Table(models.Category{}.TableName()).Delete(&existedCategory)

	if result.Error != nil {
		return nil, result.Error
	}

	return existedCategory, nil
}
