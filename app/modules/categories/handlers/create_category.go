package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *CategoryHandlers) CreateCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.CreateCategoryInput{}

		
	}
}
