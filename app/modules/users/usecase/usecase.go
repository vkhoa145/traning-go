package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/models"
	userRepo "github.com/vkhoa145/go-training/app/modules/users/repositories"
)

type UseCase interface {
	SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.UserResponse, error)
	SignInUser(ctx *fiber.Ctx, payload *models.SignInInput) (*models.UserResponse, error)
}

type UserUseCase struct {
	userRepo userRepo.UserRepoInterface
}

func NewUserUseCase(userRepo userRepo.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}
