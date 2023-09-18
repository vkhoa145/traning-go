package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (book BookUseCase) DeleteBook(ctx *fiber.Ctx, existedBook *models.Book) (*models.BookResponse, error) {
	_, err := book.bookRepo.DeleteBook(existedBook)

	if err != nil {
		return nil, err
	}

	return models.FilterBookRecord(existedBook), nil
}
