package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (cat CategoryUseCase) DeleteCategory(ctx *fiber.Ctx, existedCategory *models.Category) (*models.CategoryResponse, error) {
	_, err := cat.categoryRepo.DeleteCategory(existedCategory)

	if err != nil {
		return nil, err
	}

	return models.FilterCategoryRecord(existedCategory), nil
}
