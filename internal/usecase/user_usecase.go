package usecase

import "github.com/NikenCarolina/flashcard-be/internal/repository"

type UserUseCase interface {
	UserFlashcardUseCase
}

type userUseCase struct {
	store repository.Store
}

func NewUserUseCase(store repository.Store) UserUseCase {
	return &userUseCase{store: store}
}
