package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *BookHandlers) GetBookById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		book_id := ctx.Params("id")
		bookID, err := strconv.ParseFloat(book_id, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}
		existedBook, err := h.bookRepo.GetBookById(bookID)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterBookRecord(existedBook), "error": nil})
	}
}
