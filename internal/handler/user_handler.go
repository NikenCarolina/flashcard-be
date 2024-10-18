package handler

import "github.com/NikenCarolina/flashcard-be/internal/usecase"

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(userUserCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: userUserCase,
	}
}
