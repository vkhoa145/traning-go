package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-training/app/models"
)

func (h *UserHandlers) GetUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		user1, err := h.UserRepo.GetUserByEmail(email)
		fmt.Println(user1)

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"status": http.StatusNotFound, "error": errors.New("user not found")})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterUserRecord(user1), "error": nil})
	}
}
