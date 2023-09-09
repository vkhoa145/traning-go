package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/vkhoa145/go-training/app/models"
)

func (u UserUseCase) GetUser(ctx *fiber.Ctx, email string) (*models.UserResponse, error) {

	user, err := u.userRepo.GetUserByEmail(email)

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(user), nil
}
