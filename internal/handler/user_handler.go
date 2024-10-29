package handler

import (
	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/usecase"
)

type UserHandler struct {
	useCase usecase.UserUseCase
	config  *config.Config
}

func NewUserHandler(userUserCase usecase.UserUseCase, config *config.Config) *UserHandler {
	return &UserHandler{
		useCase: userUserCase,
		config:  config,
	}
}
