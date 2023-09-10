package handlers

import (
	bookRepo "github.com/vkhoa145/go-training/app/modules/books/repositories"
	bookUseCase "github.com/vkhoa145/go-training/app/modules/books/usecase"
)

type BookHandlers struct {
	bookRepo    bookRepo.BookRepoInterface
	bookUseCase bookUseCase.BookUseCaseInterface
}

func NewBookHandlers(bookUseCase bookUseCase.BookUseCaseInterface, bookRepo bookRepo.BookRepoInterface) *BookHandlers {
	return &BookHandlers{
		bookUseCase: bookUseCase,
		bookRepo:    bookRepo,
	}
}
