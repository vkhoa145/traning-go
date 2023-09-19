package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *BookHandlers) UpdateBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.UpdateBookInput{}
		if payloadErr := ctx.BodyParser(&payload); payloadErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": payloadErr.Error()})
		}

		book_id := ctx.Params("book_id")
		bookID, bookErr := strconv.ParseFloat(book_id, 64)
		if bookErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		category_id := ctx.Params("id")
		categoryID, catErr := strconv.ParseFloat(category_id, 64)
		if catErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}
		payload.CategoryId = uint(categoryID)

		userId := ctx.Get("User_id")
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}
		Id := uint(userIdFloat)

		existedBook, existedBookErr := h.bookRepo.GetBookById(bookID)
		if existedBookErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": existedBookErr.Error()})
		}

		if existedBook.UserId != Id {
			ctx.Status(http.StatusForbidden)
			return ctx.JSON(&fiber.Map{"status": http.StatusForbidden, "error": "Unauthorized"})
		}

		updateBook, err := h.bookUseCase.UpdateBook(ctx, &payload, existedBook)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": updateBook, "error": nil})
	}
}
