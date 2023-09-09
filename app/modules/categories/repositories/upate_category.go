package repository

import "github.com/vkhoa145/go-training/app/models"

func (r CategoryRepo) UpdateCategory(data *models.UpdateCategoryInput, existedCategory *models.Category) (*models.Category, error) {
	category := &models.Category{
		Name:        data.Name,
		Description: data.Description,
		UserId:      data.UserId,
	}

	result := r.DB.Table(models.Category{}.TableName()).Model(&existedCategory).Updates(category)

	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
