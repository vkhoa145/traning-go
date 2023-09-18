package repository

import "github.com/vkhoa145/go-training/app/models"

func (r BookRepo) DeleteBook(existedBook *models.Book) (*models.Book, error) {

	result := r.DB.Table(models.Book{}.TableName()).Delete(&existedBook)

	if result.Error != nil {
		return nil, result.Error
	}

	return existedBook, nil
}
