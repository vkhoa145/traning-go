package handlers

import (
	userRepo "github.com/vkhoa145/go-training/app/modules/users/repositories"
	user "github.com/vkhoa145/go-training/app/modules/users/usecase"
)

type UserHandlers struct {
	userUseCase user.UseCase
	UserRepo    userRepo.UserRepoInterface
}

func NewUserHandlers(userUseCase user.UseCase, UserRepo userRepo.UserRepoInterface) *UserHandlers {
	return &UserHandlers{
		userUseCase: userUseCase,
		UserRepo:    UserRepo,
	}
}
