package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (cat CategoryUseCase) CreateCategory(ctx *fiber.Ctx, payload *models.CreateCategoryInput, userId uint) (*models.CategoryResponse, error) {
	if payload.Name == "" {
		return nil, errors.New("name cant be blank")
	}

	createdCategory, err := cat.categoryRepo.CreateCategory(payload, userId)

	if err != nil {
		return nil, err
	}

	return models.FilterCategoryRecord(createdCategory), nil
}
