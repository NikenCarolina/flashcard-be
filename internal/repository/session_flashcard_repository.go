package repository

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
)

type SessionFlashcardRepository interface {
	Create(ctx context.Context, sessionID, cardID int64) error
	CheckExistsById(ctx context.Context, sessionID int) (map[int]bool, error)
	UpdateIsCorrect(ctx context.Context, sessionID, cardID int, isCorrect bool) error
}

type sessionFlashcardRepository struct {
	db database
}

func NewSessionFlashcardRepository(db database) SessionFlashcardRepository {
	return &sessionFlashcardRepository{db}
}

func (r *sessionFlashcardRepository) Create(ctx context.Context, sessionID, cardID int64) error {
	query := `
		INSERT INTO "session_flashcards"
			("session_id", "flashcard_id")
		VALUES
			($1, $2)
	`
	if _, err := r.db.ExecContext(ctx, query, sessionID, cardID); err != nil {
		return apperror.ErrInternalServerError
	}

	return nil
}

func (r *sessionFlashcardRepository) UpdateIsCorrect(ctx context.Context, sessionID, cardID int, isCorrect bool) error {
	query := `
		UPDATE "session_flashcards"
		SET 
			"is_correct" = $1,
			"is_reviewed" = TRUE
		WHERE "session_id" = $2 AND "flashcard_id" = $3
	`

	if _, err := r.db.ExecContext(ctx, query, isCorrect, sessionID, cardID); err != nil {
		return apperror.ErrInternalServerError
	}
	return nil
}

func (r *sessionFlashcardRepository) CheckExistsById(ctx context.Context, sessionID int) (map[int]bool, error) {
	query := `
		SELECT 
			"flashcard_id"
		FROM
			"session_flashcards"
		WHERE "session_id" = $1
	`

	rows, err := r.db.QueryContext(ctx, query, sessionID)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()

	flashcardIDs := map[int]bool{}
	for rows.Next() {
		var flashcardID int
		if err := rows.Scan(
			&flashcardID,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		flashcardIDs[flashcardID] = true
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return flashcardIDs, nil
}
