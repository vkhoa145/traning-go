package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/middlewares"
	"github.com/vkhoa145/go-training/app/models"
	"github.com/vkhoa145/go-training/config"
)

func (h *UserHandlers) SignInUser(config *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignInInput{}

		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		user, err := h.userUseCase.SignInUser(ctx, &payload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		t, err := middlewares.CreateAccessToken(user, config.SIGNED_STRING, 24)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		// ctx.Request().Header.Set("Authorization", t)
		ctx.Set("Authorization", t)
		ctx.Cookie(&fiber.Cookie{
			Name:    "access_token",
			Value:   t,
			Expires: time.Now().Add(time.Hour * 1),
		})
		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"status": http.StatusOK, "token": t, "error": nil})
	}
}
