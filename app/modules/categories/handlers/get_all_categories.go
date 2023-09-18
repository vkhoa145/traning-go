package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *CategoryHandlers) GetAllCategories() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		userId := uint(id)

		categories, err := h.categoryRepo.GetAllCategories(float64(userId))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": categories, "error": nil})
		// return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": "succed", "error": nil})
	}
}
