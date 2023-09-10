package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (cat CategoryUseCase) UpdateCategory(ctx *fiber.Ctx, payload *models.UpdateCategoryInput, existedCategory *models.Category) (*models.CategoryResponse, error) {
	_, err := cat.categoryRepo.UpdateCategory(payload, existedCategory)

	if err != nil {
		return nil, err
	}

	return models.FilterCategoryRecord(existedCategory), nil
}
