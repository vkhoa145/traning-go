package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *CategoryHandlers) GetCategoryById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category_id := ctx.Params("id")
		categoryID, err := strconv.ParseFloat(category_id, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		userId := uint(id)

		existedCategory, err := h.categoryRepo.GetCategoryById(categoryID, float64(userId))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterCategoryRecord(existedCategory), "error": nil})
	}
}
