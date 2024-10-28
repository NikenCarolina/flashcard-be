package usecase

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/NikenCarolina/flashcard-be/internal/model"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
)

type UserFlashcardUseCase interface {
	GetSets(ctx context.Context, userID int) ([]dto.FlashcardSet, error)
	GetSetById(ctx context.Context, userID int, setID int) (*dto.FlashcardSet, error)
	GetCards(ctx context.Context, setID int, userID int) ([]dto.Flashcard, error)
	CreateCard(ctx context.Context, userID int, setID int) (*dto.Flashcard, error)
	UpdateCard(ctx context.Context, userID int, req *dto.Flashcard) error
	DeleteCard(ctx context.Context, userID, setID, cardID int) error
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

func (u *userUseCase) GetSetById(ctx context.Context, userID int, setID int) (*dto.FlashcardSet, error) {
	flashcardSetRepo := u.store.FlashcardSet()
	flashcardSets, err := flashcardSetRepo.GetById(ctx, userID, setID)
	if err != nil {
		return nil, err
	}
	return flashcardSets.ToDto(), nil
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

func (u *userUseCase) CreateCard(ctx context.Context, userID int, setID int) (*dto.Flashcard, error) {
	res, err := u.store.Atomic(ctx, func(s repository.Store) (any, error) {
		flashcardSetRepo := u.store.FlashcardSet()
		exists, err := flashcardSetRepo.CheckExists(ctx, userID, setID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, apperror.ErrNotFound
		}

		flashcardRepo := u.store.Flashcard()
		res, err := flashcardRepo.Create(ctx, setID)
		if err != nil {
			return nil, err
		}

		progressRepo := u.store.FlashcardProgress()
		err = progressRepo.Create(ctx,
			setID,
			res.FlashcardID,
			u.flashcardConfig.RepetitionNumber,
			u.flashcardConfig.EasinessFactor,
			u.flashcardConfig.Interval,
		)
		if err != nil {
			return nil, err
		}
		return res.ToDto(), nil
	})
	if err != nil {
		return nil, err
	}
	card := res.(*dto.Flashcard)

	return card, nil
}

func (u *userUseCase) UpdateCard(ctx context.Context, userID int, req *dto.Flashcard) error {
	flashcardSetRepo := u.store.FlashcardSet()
	exists, err := flashcardSetRepo.CheckExists(ctx, userID, int(req.FlashcardSetID))
	if err != nil {
		return err
	}
	if !exists {
		return apperror.ErrNotFound
	}

	var flashcard model.Flashcard
	flashcard.LoadFromDto(*req)

	flashcardRepo := u.store.Flashcard()
	if err = flashcardRepo.Update(ctx, flashcard); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) DeleteCard(ctx context.Context, userID, setID, cardID int) error {
	flashcardSetRepo := u.store.FlashcardSet()
	exists, err := flashcardSetRepo.CheckExists(ctx, userID, setID)
	if err != nil {
		return err
	}
	if !exists {
		return apperror.ErrNotFound
	}

	flashcardRepo := u.store.Flashcard()
	if err = flashcardRepo.Delete(ctx, setID, cardID); err != nil {
		return err
	}

	return nil
}
