package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *BookHandlers) DeleteBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		book_id := ctx.Params("id")
		bookID, bookErr := strconv.ParseFloat(book_id, 64)
		if bookErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		userId := uint(id)

		existedBook, existedBookErr := h.bookRepo.GetBookById(bookID)
		if existedBookErr != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": existedBookErr.Error()})
		}

		if existedBook.UserId != userId {
			ctx.Status(http.StatusForbidden)
			return ctx.JSON(&fiber.Map{"status": http.StatusForbidden, "error": "Unauthorized"})
		}

		deleteBook, err := h.bookUseCase.DeleteBook(ctx, existedBook)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": deleteBook, "error": nil})
	}
}
