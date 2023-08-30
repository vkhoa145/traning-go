package handlers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-training/app/models"
	"github.com/vkhoa145/go-training/config"
)

func (h *UserHandlers) SignUpUser(config *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignUpInput{}

		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			errors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				A := err.Field()
				errors[A] = err.Tag()
			}

			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{"status": http.StatusUnprocessableEntity, "message": "Unprocessable Content", "errors": errors})
		}

		createdUser, err := h.userUseCase.SignUpUser(ctx, &payload)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		claims := jwt.MapClaims{
			"id":    createdUser.ID,
			"email": createdUser.Email,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(config.SIGNED_STRING))
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "token": t, "error": nil})
	}
}
