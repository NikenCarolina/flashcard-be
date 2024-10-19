package repository

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type FlashcardSetRepository interface {
	GetSets(ctx context.Context, userID int) ([]model.FlashcardSet, error)
}

type flashcardSetRepository struct {
	db database
}

func NewFlashcardSetRepository(db database) FlashcardSetRepository {
	return &flashcardSetRepository{db}
}

func (r *flashcardSetRepository) GetSets(ctx context.Context, userID int) ([]model.FlashcardSet, error) {
	flashcardSets := []model.FlashcardSet{}
	query := `
		SELECT 
			"flashcard_set_id", 
			"title", 
			"description" 
		FROM "flashcard_sets" WHERE user_id = $1`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var flashcardSet model.FlashcardSet
		if err := rows.Scan(
			&flashcardSet.FlashcardSetID,
			&flashcardSet.Title,
			&flashcardSet.Description,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		flashcardSets = append(flashcardSets, flashcardSet)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return flashcardSets, nil
}
