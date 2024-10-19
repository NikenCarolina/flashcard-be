package usecase

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/dto"
)

type UserFlashcardUseCase interface {
	GetSets(ctx context.Context, userID int) ([]dto.FlashcardSet, error)
}

func (u *userUseCase) GetSets(ctx context.Context, userID int) ([]dto.FlashcardSet, error) {
	flashcardSetRepo := u.store.FlashcardSet()
	flashcardSets, err := flashcardSetRepo.GetSets(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := []dto.FlashcardSet{}
	for _, flashcardSet := range flashcardSets {
		res = append(res, *flashcardSet.ToDto())
	}

	return res, nil
}
