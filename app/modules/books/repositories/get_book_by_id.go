package repository

import "github.com/vkhoa145/go-training/app/models"

func (r BookRepo) GetBookById(id float64) (*models.Book, error) {

	var book *models.Book
	result := r.DB.Table(models.Book{}.TableName()).Where("id=?", id).First(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}
