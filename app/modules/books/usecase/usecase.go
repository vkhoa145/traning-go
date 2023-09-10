package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
	bookRepo "github.com/vkhoa145/go-training/app/modules/books/repositories"
)

type BookUseCaseInterface interface {
	CreateBook(ctx *fiber.Ctx, payload *models.CreateBookInput) (*models.BookResponse, error)
	UpdateBook(ctx *fiber.Ctx, payload *models.UpdateBookInput, existedBook *models.Book) (*models.BookResponse, error)
	DeleteBook(ctx *fiber.Ctx, existedBook *models.Book) (*models.BookResponse, error)
}

type BookUseCase struct {
	bookRepo bookRepo.BookRepoInterface
}

func NewBookUseCase(bookRepo bookRepo.BookRepoInterface) BookUseCaseInterface {
	return &BookUseCase{bookRepo: bookRepo}
}
