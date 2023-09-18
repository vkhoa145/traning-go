package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *CategoryHandlers) GetAllCategories() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Get("User_id")
		fmt.Println("user from header", userId)
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		categories, err := h.categoryRepo.GetAllCategories(userIdFloat)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": categories, "error": nil})
		// return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": "succed", "error": nil})
	}
}
