package repository

import (
	"context"
	"fmt"

	"github.com/NikenCarolina/flashcard-be/internal/apperror"
	"github.com/NikenCarolina/flashcard-be/internal/model"
)

type FlashcardProgressRepository interface {
	Create(ctx context.Context, setID, cardID int, repetitionNumber int64, easinessFactor float64, interval int64) error
	Update(ctx context.Context, setID int, cardID int, repetitionNumber int64, easinessFactor float64, interval int64) error
	GetBySetId(ctx context.Context, setID int, limit int) ([]model.FlashcardProgress, error)
	GetById(ctx context.Context, cardID int) (*model.FlashcardProgress, error)
}

type flashcardProgressRepository struct {
	db database
}

func NewFlashcardProgressRepository(db database) FlashcardProgressRepository {
	return &flashcardProgressRepository{db}
}

func (r *flashcardProgressRepository) Create(ctx context.Context, setID, cardID int, repetitionNumber int64, easinessFactor float64, interval int64) error {
	query := `
		INSERT INTO
			"flashcard_progress"
			("flashcard_set_id","flashcard_id", "last_review", "repetition_number", "easiness_factor", "interval")
		VALUES
			($1, $2, NULL, $3, $4, $5)
	`

	if _, err := r.db.ExecContext(ctx, query,
		setID,
		cardID,
		repetitionNumber,
		easinessFactor,
		interval,
	); err != nil {
		return apperror.ErrInternalServerError
	}

	return nil
}

func (r *flashcardProgressRepository) Update(ctx context.Context, setID int, cardID int, repetitionNumber int64, easinessFactor float64, interval int64) error {
	query := fmt.Sprintf(`
		UPDATE "flashcard_progress"	
		SET 
			"last_review" = NOW(), 
			"repetition_number" = $1, 
			"easiness_factor" = $2, 
			"interval" = $3,
			"due_date" = NOW() + INTERVAL '%d DAYS' 
		WHERE
			"flashcard_id" = $4 AND "flashcard_set_id" = $5
	`, interval)
	if _, err := r.db.ExecContext(ctx, query,
		repetitionNumber,
		easinessFactor,
		interval,
		cardID,
		setID,
	); err != nil {
		return apperror.ErrInternalServerError
	}

	return nil
}

func (r *flashcardProgressRepository) GetBySetId(ctx context.Context, setID int, limit int) ([]model.FlashcardProgress, error) {
	flashcardProgresses := []model.FlashcardProgress{}
	query := `
		SELECT 
			"flashcard_id", 
			"flashcard_set_id", 
			"repetition_number",
			"easiness_factor",
			"interval",
			"last_review",
			"due_date"
		FROM "flashcard_progress" 
		WHERE flashcard_set_id = $1 
		ORDER BY due_date
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, setID, limit)
	if err != nil {
		return nil, apperror.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var flashcardProgress model.FlashcardProgress
		if err := rows.Scan(
			&flashcardProgress.FlashcardID,
			&flashcardProgress.FlashcardSetID,
			&flashcardProgress.RepetitionNumber,
			&flashcardProgress.EasinessFactor,
			&flashcardProgress.Interval,
			&flashcardProgress.LastReview,
			&flashcardProgress.DueDate,
		); err != nil {
			return nil, apperror.ErrInternalServerError
		}
		flashcardProgresses = append(flashcardProgresses, flashcardProgress)
	}

	if err := rows.Err(); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return flashcardProgresses, nil

}

func (r *flashcardProgressRepository) GetById(ctx context.Context, cardID int) (*model.FlashcardProgress, error) {
	query := `
		SELECT 
			"flashcard_id", 
			"flashcard_set_id", 
			"repetition_number",
			"easiness_factor",
			"interval",
			"last_review",
			"due_date"
		FROM "flashcard_progress" 
		WHERE "flashcard_id" = $1 
	`

	var flashcardProgress model.FlashcardProgress
	if err := r.db.QueryRowContext(ctx, query, cardID).Scan(
		&flashcardProgress.FlashcardID,
		&flashcardProgress.FlashcardSetID,
		&flashcardProgress.RepetitionNumber,
		&flashcardProgress.EasinessFactor,
		&flashcardProgress.Interval,
		&flashcardProgress.LastReview,
		&flashcardProgress.DueDate,
	); err != nil {
		return nil, apperror.ErrInternalServerError
	}

	return &flashcardProgress, nil
}
