package repository

import (
	"github.com/vkhoa145/go-training/app/models"
	"gorm.io/gorm"
)

type BookRepoInterface interface {
	CreateBook(data *models.CreateBookInput) (*models.Book, error)
	GetBookById(id float64) (*models.Book, error)
	UpdateBook(data *models.UpdateBookInput, existedBook *models.Book) (*models.Book, error)
	DeleteBook(existedBook *models.Book) (*models.Book, error)
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{DB: db}
}
