package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
	categoryRepo "github.com/vkhoa145/go-training/app/modules/categories/repositories"
)

type CategoryUseCaseInterface interface {
	CreateCategory(ctx *fiber.Ctx, payload *models.CreateCategoryInput, userId uint) (*models.CategoryResponse, error)
}

type CategoryUseCase struct {
	categoryRepo categoryRepo.CategoryRepoInterface
}

func NewCategoryUseCase(catRepo categoryRepo.CategoryRepoInterface) CategoryUseCaseInterface {
	return &CategoryUseCase{categoryRepo: catRepo}
}
