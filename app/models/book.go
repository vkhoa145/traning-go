package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	PublicDate  time.Time `gorm:"type:time" json:"time"`
	UserId      uint
	CategoryId  uint
}

func (Book) TableName() string {
	return "books"
}

type CreateBookInput struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PublicDate  time.Time `json:"public_date" validate: "required"`
	UserId      uint      `json:"user_id" validate: "required"`
	CategoryId  uint
}

type UpdateBookInput struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PublicDate  time.Time `json:"public_date" validate: "required"`
	UserId      uint      `json:"user_id" validates: "required"`
	CategoryId  uint      `json:"category_id" validate: "required"`
}

type BookResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PublicDate  time.Time `json:"public_date"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
}

func FilterBookRecord(book *Book) *BookResponse {
	return &BookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Description: book.Description,
		PublicDate:  book.PublicDate,
		UserId:      book.UserId,
		CategoryId:  book.CategoryId,
	}
}

var Books []Book
