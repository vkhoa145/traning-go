package handlers

import (
	categoryRepo "github.com/vkhoa145/go-training/app/modules/categories/repositories"
	categoryUseCase "github.com/vkhoa145/go-training/app/modules/categories/usecase"
)

type CategoryHandlers struct {
	categoryUseCase categoryUseCase.CategoryUseCaseInterface
	categoryRepo    categoryRepo.CategoryRepoInterface
}

func NewCategoryHandlers(catUseCase categoryUseCase.CategoryUseCaseInterface, catRepo categoryRepo.CategoryRepoInterface) *CategoryHandlers {
	return &CategoryHandlers{
		categoryUseCase: catUseCase,
		categoryRepo:    catRepo,
	}
}
