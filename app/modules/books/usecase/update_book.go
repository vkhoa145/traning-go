package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (book BookUseCase) UpdateBook(ctx *fiber.Ctx, payload *models.UpdateBookInput, existedBook *models.Book) (*models.BookResponse, error) {
	_, err := book.bookRepo.UpdateBook(payload, existedBook)

	if err != nil {
		return nil, err
	}

	return models.FilterBookRecord(existedBook), nil
}
