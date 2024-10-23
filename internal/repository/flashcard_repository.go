package repository

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type FlashcardRepository interface {
	GetBySetId(ctx context.Context, setID int) ([]model.Flashcard, error)
}

type flashcardRepository struct {
	db database
}

func NewFlashcardRepository(db database) FlashcardRepository {
	return &flashcardRepository{db}
}

func (r *flashcardRepository) GetBySetId(ctx context.Context, setID int) ([]model.Flashcard, error) {
	flashcards := []model.Flashcard{}
	query := `
		SELECT 
			"flashcard_id", 
			"flashcard_set_id", 
			"term", 
			"definition"
		FROM flashcards
		WHERE "flashcard_set_id" = $1 
	`

	rows, err := r.db.QueryContext(ctx, query, setID)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var flashcard model.Flashcard
		if err := rows.Scan(
			&flashcard.FlashcardID,
			&flashcard.FlashcardSetID,
			&flashcard.Term,
			&flashcard.Definition,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		flashcards = append(flashcards, flashcard)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return flashcards, nil
}
