package handlers

import (
	categoryUseCase "github.com/vkhoa145/go-training/app/modules/categories/usecase"
)

type CategoryHandlers struct {
	categoryUseCase categoryUseCase.CategoryUseCaseInterface
}

func NewCategoryHandlers(catUseCase categoryUseCase.CategoryUseCaseInterface) *CategoryHandlers {
	return &CategoryHandlers{
		categoryUseCase: catUseCase,
	}
}
