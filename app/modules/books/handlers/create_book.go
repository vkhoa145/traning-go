package handlers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *BookHandlers) CreateBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.CreateBookInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		category_id := ctx.Params("id")
		categoryID, catErr := strconv.ParseFloat(category_id, 64)
		if catErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
		}

		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			errors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				A := err.Field()
				field, _ := reflect.TypeOf(payload).FieldByName(A)
				jsonTag := field.Tag.Get("json")
				errors[jsonTag] = err.Tag()
			}

			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{"status": http.StatusUnprocessableEntity, "message": "Unprocessable Content", "errors": errors})
		}

		userId := ctx.Get("User_id")
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}
		payload.UserId = uint(userIdFloat)
		payload.CategoryId = uint(categoryID)

		createdBook, err := h.bookUseCase.CreateBook(ctx, &payload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": createdBook, "error": nil})
	}
}
