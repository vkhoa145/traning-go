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

		// why cannot get by email from claims
		email := claims["email"].(string)
		// why id must be float64
		id := claims["id"].(float64)
		fmt.Println("user email", email)
		// user1, err := h.UserRepo.GetUserByEmail(email)
		user2, err := h.UserRepo.GetUserById(id)

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"status": http.StatusNotFound, "error": errors.New("user not found")})
		}

		ctx.Status(http.StatusCreated)
		// return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterUserRecord(user1), "error": nil})
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": models.FilterUserRecord(user2), "error": nil})
	}
}
