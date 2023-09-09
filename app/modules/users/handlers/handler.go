package handlers

import (
	user "github.com/vkhoa145/go-training/app/modules/users/usecase"
)

type UserHandlers struct {
	userUseCase user.UseCase
}

func NewUserHandlers(userUseCase user.UseCase) *UserHandlers {
	return &UserHandlers{
		userUseCase: userUseCase,
	}
}
