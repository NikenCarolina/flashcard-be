package usecase

import (
	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
)

type UserUseCase interface {
	UserFlashcardUseCase
	UserSessionUseCase
}

type userUseCase struct {
	store           repository.Store
	flashcardConfig config.FlashcardConfig
}

func NewUserUseCase(store repository.Store, flashcardConfig config.FlashcardConfig) UserUseCase {
	return &userUseCase{store: store, flashcardConfig: flashcardConfig}
}
