package usecase

import (
	"github.com/NikenCarolina/flashcard-be/internal/config"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
	"github.com/NikenCarolina/flashcard-be/internal/util"
)

type UserUseCase interface {
	UserFlashcardUseCase
	UserSessionUseCase
	UserAuthUseCase
}

type userUseCase struct {
	store           repository.Store
	flashcardConfig config.FlashcardConfig
	jwtProvider     util.JwtProvider
	bycryptProvider util.BycryptProvider
}

func NewUserUseCase(store repository.Store, flashcardConfig config.FlashcardConfig, jwtProvider util.JwtProvider, bycryptProvider util.BycryptProvider) UserUseCase {
	return &userUseCase{store: store, flashcardConfig: flashcardConfig, jwtProvider: jwtProvider, bycryptProvider: bycryptProvider}
}
