package usecase

import "github.com/NikenCarolina/flashcard-be/internal/repository"

type UserUseCase interface{}

type userUseCase struct {
	store repository.Store
}

func NewUserUseCase(store repository.Store) *userUseCase {
	return &userUseCase{store: store}
}
