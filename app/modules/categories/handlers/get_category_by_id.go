package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *CategoryHandlers) GetCategoryById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category_id := ctx.Params("id")
		categoryID, err := strconv.ParseFloat(category_id, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		userId := ctx.Get("User_id")
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		existedCategory, err := h.categoryRepo.GetCategoryById(categoryID, userIdFloat)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterCategoryRecord(existedCategory), "error": nil})
	}
}
