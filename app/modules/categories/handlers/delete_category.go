package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *CategoryHandlers) DeleteCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category_id := ctx.Params("id")
		categoryID, catErr := strconv.ParseFloat(category_id, 64)
		if catErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		userId := ctx.Get("User_id")
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		existedCategory, existedCategoryErr := h.categoryRepo.GetCategoryById(categoryID, userIdFloat)
		if existedCategoryErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": existedCategoryErr.Error()})
		}

		deleteCategory, err := h.categoryUseCase.DeleteCategory(ctx, existedCategory)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": deleteCategory, "error": nil})
	}
}
