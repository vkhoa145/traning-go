package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *CategoryHandlers) UpdateCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.UpdateCategoryInput{}
		if payloadErr := ctx.BodyParser(&payload); payloadErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": payloadErr.Error()})
		}

		category_id := ctx.Params("id")
		categoryID, catErr := strconv.ParseFloat(category_id, 64)
		if catErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		userId := uint(id)

		existedCategory, existedCategoryErr := h.categoryRepo.GetCategoryById(categoryID, float64(userId))
		if existedCategoryErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": existedCategoryErr.Error()})
		}

		updateCategory, err := h.categoryUseCase.UpdateCategory(ctx, &payload, existedCategory)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": updateCategory, "error": nil})
	}
}
