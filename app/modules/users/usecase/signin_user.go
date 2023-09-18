package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/vkhoa145/go-training/app/models"
)

func (u UserUseCase) SignInUser(ctx *fiber.Ctx, payload *models.SignInInput) (*models.UserResponse, error) {
	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(user), nil
}
