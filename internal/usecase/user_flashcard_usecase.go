package usecase

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
)

type UserFlashcardUseCase interface {
	GetSets(ctx context.Context, userID int) ([]dto.FlashcardSet, error)
	GetCards(ctx context.Context, setID int, userID int) ([]dto.Flashcard, error)
}

func (u *userUseCase) GetSets(ctx context.Context, userID int) ([]dto.FlashcardSet, error) {
	flashcardSetRepo := u.store.FlashcardSet()
	flashcardSets, err := flashcardSetRepo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := []dto.FlashcardSet{}
	for _, flashcardSet := range flashcardSets {
		res = append(res, *flashcardSet.ToDto())
	}

	return res, nil
}

func (u *userUseCase) GetCards(ctx context.Context, userID int, setID int) ([]dto.Flashcard, error) {
	flashcardSetRepo := u.store.FlashcardSet()
	exists, err := flashcardSetRepo.CheckExists(ctx, userID, setID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, apperror.ErrNotFound
	}

	flashcardRepo := u.store.Flashcard()
	flashcards, err := flashcardRepo.GetBySetId(ctx, setID)
	if err != nil {
		return nil, err
	}

	res := []dto.Flashcard{}
	for _, flashcard := range flashcards {
		res = append(res, *flashcard.ToDto())
	}

	return res, nil
}
