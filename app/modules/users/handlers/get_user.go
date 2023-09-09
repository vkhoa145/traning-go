package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *UserHandlers) GetUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		user1, err := h.userUseCase.GetUser(ctx, email)

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"status": http.StatusNotFound, "error": errors.New("user not found")})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": user1, "error": nil})
	}
}
