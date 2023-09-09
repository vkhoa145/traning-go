package handlers

import (
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *CategoryHandlers) CreateCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.CreateCategoryInput{}

		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
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

		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		payload.UserId = uint(id)

		createdCategory, err := h.categoryUseCase.CreateCategory(ctx, &payload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": createdCategory, "error": nil})
	}
}
