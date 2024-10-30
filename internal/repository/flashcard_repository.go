package repository

import (
	"context"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type FlashcardRepository interface {
	GetByCardId(ctx context.Context, cardID int) (*model.Flashcard, error)
	GetBySetId(ctx context.Context, setID int) ([]model.Flashcard, error)
	Create(ctx context.Context, setID int) (*model.Flashcard, error)
	Update(ctx context.Context, flashcard model.Flashcard) error
	Delete(ctx context.Context, setID, cardID int) error
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
		ORDER BY "created_at"
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

func (r *flashcardRepository) GetByCardId(ctx context.Context, cardID int) (*model.Flashcard, error) {
	query := `
		SELECT 
			"flashcard_id", 
			"flashcard_set_id", 
			"term", 
			"definition"
		FROM flashcards
		WHERE "flashcard_id" = $1 
	`

	var flashcard model.Flashcard
	err := r.db.QueryRowContext(ctx, query, cardID).Scan(
		&flashcard.FlashcardID,
		&flashcard.FlashcardSetID,
		&flashcard.Term,
		&flashcard.Definition,
	)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return &flashcard, nil
}

func (r *flashcardRepository) Create(ctx context.Context, setID int) (*model.Flashcard, error) {
	query := `
		INSERT INTO 
			"flashcards"
			("flashcard_set_id", "created_at", "updated_at")
		VALUES
			($1, NOW(), NOW())
		RETURNING
			"flashcard_id", "flashcard_set_id", "term", "definition"
	`

	var res model.Flashcard
	if err := r.db.QueryRowContext(ctx, query, setID).Scan(
		&res.FlashcardID,
		&res.FlashcardSetID,
		&res.Term,
		&res.Definition,
	); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return &res, nil
}

func (r *flashcardRepository) Update(ctx context.Context, flashcard model.Flashcard) error {
	query := `
		UPDATE "flashcards"
		SET "term" = $1, "definition" = $2
		WHERE "flashcard_set_id" = $3 AND "flashcard_id" = $4 
	`
	if _, err := r.db.ExecContext(ctx, query,
		flashcard.Term,
		flashcard.Definition,
		flashcard.FlashcardSetID,
		flashcard.FlashcardID,
	); err != nil {
		return apperror.ErrInternalServerError
	}
	return nil
}

func (r *flashcardRepository) Delete(ctx context.Context, setID, cardID int) error {
	query := `
		DELETE FROM "flashcards"
		WHERE "flashcard_set_id" = $1 AND "flashcard_id" = $2 
	`
	res, err := r.db.ExecContext(ctx, query, setID, cardID)
	if err != nil {
		return apperror.ErrInternalServerError
	}

	rowNum, err := res.RowsAffected()
	if err != nil {
		return apperror.ErrInternalServerError
	}
	if rowNum == 0 {
		return apperror.ErrNotFound
	}

	return nil
}
