package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (book BookUseCase) CreateBook(ctx *fiber.Ctx, payload *models.CreateBookInput) (*models.BookResponse, error) {
	createBook, err := book.bookRepo.CreateBook(payload)

	if err != nil {
		return nil, err
	}

	return models.FilterBookRecord(createBook), nil
}
