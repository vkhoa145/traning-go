package repository

import "github.com/vkhoa145/go-training/app/models"

func (r BookRepo) CreateBook(data *models.CreateBookInput) (*models.Book, error) {
	book := &models.Book{
		Name:        data.Name,
		Description: data.Description,
		PublicDate:  data.PublicDate,
		UserId:      data.UserId,
		CategoryId:  data.CategoryId,
	}

	result := r.DB.Table(models.Book{}.TableName()).Create(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}
