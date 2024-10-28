package usecase

import (
	"context"
	"log"
	"math"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/dto"
	"github.com/NikenCarolina/flashcard-be/internal/repository"
)

type UserSessionUseCase interface {
	StartSession(ctx context.Context, userID, setID int) (*dto.Session, error)
	EndSession(ctx context.Context, userID, sessionID, setID int, req dto.EndSessionRequest) error
}

func (u *userUseCase) StartSession(ctx context.Context, userID, setID int) (*dto.Session, error) {
	res, err := u.store.Atomic(ctx, func(s repository.Store) (any, error) {
		flashcardSetRepo := s.FlashcardSet()
		exists, err := flashcardSetRepo.CheckExists(ctx, userID, setID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, apperror.ErrNotFound
		}

		log.Println("SessionRepo")

		sessionRepo := s.Session()
		session, err := sessionRepo.Create(ctx, userID, setID)
		if err != nil {
			return nil, err
		}

		log.Println("FlashcardProgress")
		flashcardProgressRepo := s.FlashcardProgress()
		flashcardProgresses, err := flashcardProgressRepo.GetBySetId(ctx, setID, 10)
		if err != nil {
			return nil, err
		}

		log.Println("SessionFlashcardRepo")
		var sessionFlashcards []dto.SessionFlashcard
		sessionCardRepo := s.SessionFlashcard()
		cardRepo := s.Flashcard()
		for _, progress := range flashcardProgresses {
			log.Println("SessionFlashcardRepo")
			err = sessionCardRepo.Create(ctx, session.SessionID, progress.FlashcardID)
			if err != nil {
				return nil, err
			}
			log.Println("CardRepo")
			card, err := cardRepo.GetByCardId(ctx, int(progress.FlashcardID))
			if err != nil {
				return nil, err
			}
			flashcardSession := &dto.SessionFlashcard{
				Flashcard:         *card.ToDto(),
				FlashcardProgress: *progress.ToDto(),
			}
			sessionFlashcards = append(sessionFlashcards, *flashcardSession)
		}

		sessionRes := dto.Session{
			SessionID:  int(session.SessionID),
			Flashcards: sessionFlashcards,
		}
		return sessionRes, nil
	})
	if err != nil {
		return nil, err
	}
	cardSession := res.(dto.Session)
	return &cardSession, err
}

func (u *userUseCase) EndSession(ctx context.Context, userID, sessionID, setID int, req dto.EndSessionRequest) error {
	_, err := u.store.Atomic(ctx, func(s repository.Store) (any, error) {
		flashcardSetRepo := s.FlashcardSet()
		exists, err := flashcardSetRepo.CheckExists(ctx, userID, setID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, apperror.ErrNotFound
		}

		sessionRepo := s.Session()
		exists, err = sessionRepo.CheckExists(ctx, sessionID, userID, setID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, apperror.ErrNotFound
		}

		sessionFlashcardRepo := s.SessionFlashcard()
		sessionFlashcardIDs, err := sessionFlashcardRepo.CheckExistsById(ctx, sessionID)
		if err != nil {
			return nil, err
		}

		flashcardProgressRepo := s.FlashcardProgress()
		for _, value := range req.Flashcards {
			if _, ok := sessionFlashcardIDs[value.FlashcardID]; !ok {
				return nil, apperror.ErrBadRequest
			}

			err := sessionFlashcardRepo.UpdateIsCorrect(ctx, sessionID, value.FlashcardID, *value.IsCorrect)
			if err != nil {
				return nil, err
			}

			progress, err := flashcardProgressRepo.GetById(ctx, value.FlashcardID)
			if err != nil {
				return nil, err
			}
			var repetitionNumber int64
			var easinessFactor float64
			var interval int64
			var grade float64
			if *value.IsCorrect {
				if progress.RepetitionNumber == 0 {
					interval = 1
				} else if progress.RepetitionNumber == 1 {
					interval = 6
				} else {
					interval = int64(math.Round(float64(progress.Interval) * progress.EasinessFactor))
				}
				grade = 5
				repetitionNumber = progress.RepetitionNumber + 1
			} else {
				repetitionNumber = 0
				interval = 1
			}
			easinessFactor = progress.EasinessFactor + (0.1 - (5-grade)*(0.08+(5-grade)*0.02))
			if easinessFactor < 1.3 {
				easinessFactor = 1.3
			}
			err = flashcardProgressRepo.Update(ctx, setID, value.FlashcardID, repetitionNumber, easinessFactor, interval)
			if err != nil {
				return nil, err
			}

		}

		err = sessionRepo.EndById(ctx, sessionID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}
